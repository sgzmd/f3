package main

import (
	"compress/gzip"
	"flag"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/sgzmd/f3/data/helpers"

	"github.com/anaskhan96/soup"
)

var BannedPrefixes = [...]string{"lib.b", "lib.a", "lib.reviews", "lib.md5"}

func CreateUrlList(baseUrl string) []string {
	var urls []string
	resp, err := soup.Get(baseUrl)
	if err != nil {
		fmt.Printf("Error downloading base URL: %s", err.Error())
		os.Exit(1)
	}
	doc := soup.HTMLParse(resp)
	links := doc.FindAll("a")
	for _, link := range links {
		href := link.Attrs()["href"]

		if !strings.HasPrefix(href, "lib") {
			continue
		}

		bannedPrefixFound := false
		for _, prefix := range BannedPrefixes {
			if strings.HasPrefix(href, prefix) {
				bannedPrefixFound = true
				break
			}
		}

		if bannedPrefixFound {
			continue
		}

		u := fmt.Sprintf("%s%s", baseUrl, href)
		fmt.Printf("%s -> %s\n", link.Text(), u)
		urls = append(urls, u)
	}
	return urls
}

func main() {
	const BASE_URL = "http://flibusta.site/sql/"
	baseUrl := flag.String("base_url", BASE_URL, "Base URL of the Flibusta SQL download page")
	download := flag.Bool("download", true, "Whether new files are to be downloaded")
	dump := flag.String("dump_file", "flibusta.sql", "Name of the dump file to be created")

	flag.Parse()

	if *download {
		log.Printf("Downloading from %s to %s", *baseUrl, *dump)
		urls := CreateUrlList(*baseUrl)

		var wg sync.WaitGroup

		gzippedFiles := make([]string, 0)

		sqldump, err := os.Create(*dump)
		if err != nil {
			log.Fatal(err)
		}

		var mu sync.Mutex
		for _, u := range urls {
			wg.Add(1)
			go func(u string) {
				defer wg.Done()

				// Detecting the name of the file and creating the temp file
				parsedUrl, err := url.Parse(u)
				if err != nil {
					log.Fatal(err)
				}
				tempFile, err := ioutil.TempFile("", filepath.Base(parsedUrl.Path))
				if err != nil {
					log.Fatal(err)
				}
				defer os.Remove(tempFile.Name())
				filename := tempFile.Name()

				log.Printf("Starting download from %s -> %s", u, filename)
				err = helpers.DownloadFile(filename, u)
				if err != nil {
					log.Fatal(err)
				}

				log.Printf("%s -> %s finished.", u, filename)
				gzippedFiles = append(gzippedFiles, filename)

				f, err := os.Open(filename)
				if err != nil {
					log.Fatal(err)
				}
				gzr, err := gzip.NewReader(f)
				if err != nil {
					log.Fatal(err)
				}
				bytes, err := ioutil.ReadAll(gzr)
				if err != nil {
					log.Fatal(err)
				}

				mu.Lock()
				defer mu.Unlock()

				n, err := sqldump.Write(bytes)
				if n != len(bytes) {
					log.Printf("Written only %d bytes out of %d to %s", n, len(bytes), *dump)
				}
				if err != nil {
					log.Fatal(err)
				}
			}(u)
		}

		wg.Wait()

		err = sqldump.Close()
		if err != nil {
			log.Fatal(err)
		}

		os.Chmod(*dump, fs.ModePerm)
	}
}
