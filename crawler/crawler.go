package crawler

import (
	"math/rand"
	"net/http"
	"golang.org/x/net/html"
	"fmt"
)

type Crawler struct {
	baseUrl string
	initalDelay int
	depth int
	vistedUrls []string
}

func NewCrawler(baseUrl string) *Crawler {
	crawler := new(Crawler)
	crawler.baseUrl = baseUrl
	crawler.initalDelay = rand.Intn(10)
	crawler.depth = rand.Intn(3)
	crawler.vistedUrls = make([]string, 0)

	return crawler
}

func (self *Crawler) Crawl() {
	url := self.baseUrl

	resp, err := http.Get(url)

	if err != nil {
		return
	}

	defer resp.Body.Close()

	tokenizer := html.NewTokenizer(resp.Body)


	for {
		next := tokenizer.Next()

		switch next {
		case html.ErrorToken:
			return
		case html.StartTagToken:
			token := tokenizer.Token()

			if token.Data == "a" {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						fmt.Println("Found href:", attr.Val)
						break
					}
				}
			}
		}
	}
}