package main

import (
	"github.com/nitishparkar/muffliato/crawler"
)

func main() {
	crawler := crawler.NewCrawler("http://nitishparkar.com/")
	crawler.Crawl()
}
