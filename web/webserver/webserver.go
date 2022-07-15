package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
	"github.com/jessevdk/go-flags"
	"github.com/sgzmd/f3/web/common"
	"github.com/sgzmd/f3/web/rpc"
	"github.com/sgzmd/f3/web/webserver/handlers"
	"github.com/sgzmd/go-telegram-auth/tgauth"
	"log"
)

var (
	auth tgauth.TelegramAuth
)

func main() {
	var opts common.Options

	_, err := flags.Parse(&opts)

	if err != nil {
		log.Fatal(err)
	}

	auth = tgauth.NewTelegramAuth(opts.TelegramToken, "/login", "/check-auth")
	client, err := rpc.NewClient(&opts.GrpcBackend)
	if err != nil {
		log.Fatal(err)
	}

	clientContext := handlers.ClientContext{
		RpcClient: client,
		Auth:      &auth,
	}

	engine := html.New("./templates/web", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Use(logger.New())

	app.Static(handlers.StaticPrefix, "./templates/static")

	app.Use(handlers.Auth(auth))

	app.Get("/", handlers.IndexHandler(clientContext))
	app.Get("/search/:searchTerm", handlers.SearchHandler(clientContext))
	app.Get("/track/:entityType/:id", handlers.TrackUntrackHandler(clientContext, handlers.Track))
	app.Get("/untrack/:entityType/:id", handlers.TrackUntrackHandler(clientContext, handlers.Untrack))
	app.Get(handlers.Login, LoginHandler())

	log.Fatal(app.Listen(":8080"))
}

func LoginHandler() func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		return ctx.Render("login", fiber.Map{})
	}
}
