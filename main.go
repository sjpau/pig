package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gocolly/colly/v2"
)

/*TODO:
Introduce wait group for downloading files.
Dirwalking
*/
func main() {
	crawler := colly.NewCollector()
	crawler.Limit(&colly.LimitRule{
		DomainGlob:  os.Args[1] + "/*",
		Delay:       1 * time.Second,
		RandomDelay: 1 * time.Second,
	})
	crawler.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Println(e.Request.AbsoluteURL(link))
	})
	crawler.Visit(os.Args[1])
}
