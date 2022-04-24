package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"

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

	if *download {
		urls := CreateUrlList(*baseUrl)

		for _, u := range urls {
			parsedUrl, err := url.Parse(u)
			if err != nil {
				panic(err)
			}

			fileName := filepath.Base(parsedUrl.Path)
			helpers.DownloadFile(fileName, u)
		}
	}
}
