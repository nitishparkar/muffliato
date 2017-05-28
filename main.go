package main

import (
	"fmt"
	"github.com/nitishparkar/muffliato/crawler"
	"io/ioutil"
	"strings"
)

const sitesFile string = "sites.txt"

func main() {
	data, err := ioutil.ReadFile(sitesFile)

	if err != nil {
		panic(err)
	}

	sites := strings.Split(string(data), "\n")

	done := make(chan bool)

	for _, site := range sites {
		go func(site string) {
			crawler := crawler.NewCrawler(site)
			crawler.Crawl()

			done <- true
		}(site)
	}

	<-done

	fmt.Println("Exiting")
}
