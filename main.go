package main

import (
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/sjpau/pig/component/grep"
)

type HASH [32]byte

func main() {
	InitCmdOptions()
	files := make(map[HASH]File)
	domain := grep.DomainNameFromURL(FlagURL)
	crawler := colly.NewCollector(
		colly.AllowedDomains(domain),
	)

	crawler.Limit(&colly.LimitRule{
		DomainGlob:  FlagURL + "/*",
		Delay:       time.Duration(FlagLazy) * time.Second,
		RandomDelay: 1 * time.Second,
	})

	crawler.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		url := e.Request.AbsoluteURL(link)
		if grep.IsFileFormat(link, FlagExt) {
			tmp := File{
				Path: url,
				Name: grep.FileNameFromPath(link),
			}
			files[sha256.Sum256([]byte(url))] = tmp
		} else {
			if FlagVerbose {
				fmt.Println("Visiting", url)
			}
			crawler.Visit(url)
		}
	})

	crawler.Visit(FlagURL)
	for _, value := range files {
		if FlagAsk {
			if value.Ask("Download") {
				value.Download(FlagDest)
			}
		} else {
			value.Download(FlagDest)
		}
	}
}
