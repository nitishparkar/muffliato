package crawler

import (
	"math/rand"
	"net/http"
	"golang.org/x/net/html"
	"fmt"
	"net/url"
	"time"
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

	toVisit := make([]url.URL, 0)

	for {
		next := tokenizer.Next()

		if next == html.ErrorToken {
			break
		}

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
							toVisit = append(toVisit, *u)
						}
					}
					break
				}
			}
		}
	}

	numConplete := 0

	for _, urlToVisit := range toVisit {
		go func(urlToVisit url.URL) {
			time.Sleep(time.Duration(rand.Intn(len(toVisit) * 2)) * time.Second)
			fmt.Println("Visitng: ", urlToVisit.String())
			_, _ = http.Get(urlToVisit.String())
			numConplete++
		}(urlToVisit)
	}

	for numConplete < len(toVisit) {
		time.Sleep(1 * time.Second)
	}
}