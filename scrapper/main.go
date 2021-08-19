package main

import (
	"strings"
	"github.com/gocolly/colly/v2"
)


func main() {

	t := &Collector{
		url: "https://diablo2build.fr/objets/",
		c:colly.NewCollector(),
	}
	t.c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		b := e.Attr("href")
		if strings.HasPrefix(b, t.url) {
			if len(b) > len(t.url) {
				t.child = append(t.child, b)
				//e.Request.Visit(b)
			}
		}
	})

	t.c.OnRequest(func(r *colly.Request) {
		//r.URL
	})
	t.c.Visit(t.url)
	child(t.child)
}
