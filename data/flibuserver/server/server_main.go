package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"github.com/dgraph-io/badger/v3"
	"github.com/sgzmd/f3/data/flibuserver/server/flibustadb"
	"github.com/sgzmd/f3/data/gen/go/flibuserver/proto/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/exec"
	"time"

	"github.com/go-sql-driver/mysql"
)

var (
	port       = flag.Int("port", 9000, "RPC server port")
	flibustaDb = flag.String("flibusta_db", "", "Path to Flibusta SQLite3 database")
	datastore  = flag.String("datastore", "", "Path to the data store to use")
	update     = flag.Duration("update_every", 24*time.Hour, "How often to re-download files")
	updateCmd  = flag.String("update_cmd", "/app/downloader_launcher.sh", "Command to kick-off re-download")

	mysqlHost = flag.String("mysql_host", "localhost", "MySQL host")
	mysqlPort = flag.String("mysql_port", "3306", "MySQL port")
	mysqlUser = flag.String("mysql_user", "root", "MySQL user")
	mysqlPass = flag.String("mysql_pass", "", "MySQL password")
	mysqlDb   = flag.String("mysql_db", "flibusta", "MySQL database")
)

// Interceptor is an unary interceptor for GRPC, logging each request before execution
func Interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("Received request: %s with request %v", info.FullMethod, req)
	h, err := handler(ctx, req)
	log.Printf("Finished request %s with resoponse %v and error %v", info.FullMethod, h, err)
	return h, err
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(Interceptor))
	srv, err := NewServer(*flibustaDb, *datastore)
	if err != nil {
		log.Fatalf("Couldn't create server: %v", err)
		os.Exit(2)
	}
	defer srv.Close()

	proto.RegisterFlibustierServiceServer(s, srv)

	reflection.Register(s)
	log.Printf("server listening at %v", lis.Addr())

	// Register GRPC healthcheck service
	healthServer := health.NewServer()
	healthServer.SetServingStatus("", grpc_health_v1.HealthCheckResponse_SERVING)
	s.RegisterService(&grpc_health_v1.Health_ServiceDesc, healthServer)

	if updateCmd != nil {
		dbReopen := time.NewTicker(*update)

		log.Printf("Scheduling to run %s every %s", *updateCmd, *update)

		go func() {
			for range dbReopen.C {
				downloadCmd := exec.Command(*updateCmd)
				downloadCmd.Stdout = os.Stdout
				downloadCmd.Stderr = os.Stderr

				err = downloadCmd.Start()
				if err != nil {
					log.Printf("Failed to download database update: %+v", err)
					continue
				}

				downloadCmd.Wait()

				log.Printf("Re-opening database ...")
				srv.Lock.Lock()
				db, err := OpenDatabase(*flibustaDb)
				srv.Lock.Unlock()
				if err != nil {
					log.Fatalf("Failed to open database: %s", err)
					os.Exit(1)
				}

				srv.db = flibustadb.NewFlibustaSqlite(db)

				log.Printf("Database re-opened.")
			}
		}()
	}

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func OpenDatabase(db_path string) (*sql.DB, error) {
	return sql.Open("sqlite3", db_path)
}

func NewServer(db_path string, datastore string) (*server, error) {
	srv := new(server)

	cfg := mysql.Config{
		User:                 *mysqlUser,
		Passwd:               *mysqlPass,
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%s", *mysqlHost, *mysqlPort),
		DBName:               *mysqlDb,
		AllowNativePasswords: true,
	}

	log.Printf("Connecting to MariaDB: %+v", cfg.FormatDSN())

	mariaDb, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}

	pingErr := mariaDb.Ping()
	if pingErr != nil {
		log.Fatalf("Failed to ping MariaDB: %+v", pingErr)
	}
	srv.db = flibustadb.NewFlibustaSqlMariaDb(mariaDb)

	var opt badger.Options
	if datastore == "" {
		opt = badger.DefaultOptions("").WithInMemory(true)
	} else {
		opt = badger.DefaultOptions(datastore)
	}

	srv.data, err = badger.Open(opt)
	if err != nil {
		return nil, err
	}

	return srv, nil
}
