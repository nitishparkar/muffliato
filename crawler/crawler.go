package crawler

import (
	"math/rand"
	"net/http"
	"golang.org/x/net/html"
	"fmt"
	"net/url"
)

type Crawler struct {
	baseUrl *url.URL
	initalDelay int
}

func NewCrawler(baseUrl string) *Crawler {
	crawler := new(Crawler)
	crawler.baseUrl, _ = url.Parse(baseUrl)
	crawler.initalDelay = rand.Intn(10)

	return crawler
}

func (self *Crawler) Crawl() {
	baseUrl := self.baseUrl

	resp, err := http.Get(baseUrl.String())

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
						u, err := url.Parse(attr.Val)
						if err == nil {
							if !u.IsAbs() {
								if u.String() == "/" {
									continue
								}
								u.Scheme = baseUrl.Scheme
								u.Host = baseUrl.Host
								fmt.Println("Visiting:", u.String())
								_, _ = http.Get(u.String())
							}
						}
						break
					}
				}
			}
		}
	}
}