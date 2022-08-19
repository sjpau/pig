package main

import (
	"context"
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/sjpau/pig/component/grep"
)

type HASH [32]byte

func CrawlWithTimeout(ctx context.Context, timeout time.Duration, crawler *colly.Collector, files map[HASH]File) {
	if timeout != 0 {
		ctx, cancel := context.WithTimeout(ctx, timeout)
		stop := false
		defer cancel()
		go func(ctx context.Context) {
			for {
				select {
				case <-ctx.Done():
					stop = true
					cancel()
					return
				}
			}
		}(ctx)
		crawler.OnRequest(func(r *colly.Request) {
			if stop {
				r.Abort()
			}
		})
	}
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
}

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
	ctx := context.Background()
	CrawlWithTimeout(ctx, time.Duration(FlagTimer)*time.Second, crawler, files)
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
