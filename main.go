package main

import (
	"crypto/sha256"
	"fmt"
	"os"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/sjpau/pig/component/grep"
)

type HASH [32]byte

/*TODO:
Lazy crawling
Extension specification
Put crawling on a gorouting?
*/
func main() {

	files := make(map[HASH]File)
	domain := grep.DomainNameFromURL(os.Args[1])
	crawler := colly.NewCollector(
		colly.AllowedDomains(domain),
	)

	crawler.Limit(&colly.LimitRule{
		DomainGlob:  os.Args[1] + "/*",
		Delay:       1 * time.Second,
		RandomDelay: 1 * time.Second,
	})

	crawler.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		url := e.Request.AbsoluteURL(link)
		if grep.IsFileFormat(link) {
			tmp := File{
				Path: url,
				Name: grep.FileNameFromPath(link),
			}
			files[sha256.Sum256([]byte(url))] = tmp
		} else {
			fmt.Println("Visiting", url)
			crawler.Visit(url)
		}
	})

	crawler.Visit(os.Args[1])
	for _, value := range files {
		if value.Ask("Download") {
			value.Download()
		}
	}
}
