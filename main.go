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

	validSites := make([]string, 0)

	for _, site := range sites {
		t := strings.TrimSpace(site)
		if t != "" {
			validSites = append(validSites, t)
		}
	}

	done := make(chan bool, len(validSites))

	for _, site := range validSites {
		go func(site string) {
			crawler := crawler.NewCrawler(site)
			crawler.Crawl()

			done <- true
		}(site)
	}

	for _ = range done {

	}

	fmt.Println("Exiting")
}
