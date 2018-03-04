package main

import (
	"github.com/MRushtonScratch/goLearn/htmlParser"
	"log"
	"fmt"
)

func main () {
	sites := []string {
		"http://www.bbc.co.uk/sport/football",
		"https://www.telegraph.co.uk/football/",
	}

	for _, url := range sites {
		node, err := htmlParser.GetHtmlNode(url)

		if err != nil {
			log.Fatal("Unable to get url %s", url)
		}

		fmt.Println(url)
		fmt.Println("-------------------------------")


		for _, link := range htmlParser.GetLinks(nil, node) {
			fmt.Println(link)
		}
	}


}
