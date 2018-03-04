package htmlParser

import (
	"golang.org/x/net/html"
	"net/http"
	"log"
)

func GetHtmlNode(url string) (*html.Node, error) {
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal("Unable to get url %s", url)
	}

	defer resp.Body.Close()

	return html.Parse(resp.Body)
}

func GetLinks(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	for c := n.FirstChild;c!=nil;c =c.NextSibling {
		links = GetLinks(links, c)
	}

	return links
}
