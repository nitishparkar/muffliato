package main

import (
	"fmt"
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

	for _, site := range(sites) {
		fmt.Println(site)
	}
}
