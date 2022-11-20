package crawler

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// Visit appends to links each link found in n and returns the result.
func Visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode &&
		(n.Data == "a" ||
			n.Data == "script" ||
			n.Data == "img" ||
			n.Data == "style") {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	if c := n.FirstChild; c != nil {
		links = Visit(links, c)
	}
	if n.NextSibling != nil {
		links = Visit(links, n.NextSibling)
	}
	return links
}

// CountElements count elements
func CountElements(countEl map[string]int, n *html.Node) map[string]int {
	_, isMapContainsKey := countEl[n.Data]
	if n.Type == html.ElementNode && isMapContainsKey {
		countEl[n.Data]++
	}
	if n.Type == html.ElementNode && !isMapContainsKey {
		countEl[n.Data] = 1
	}

	if c := n.FirstChild; c != nil {
		countEl = CountElements(countEl, c)
	}
	if n.NextSibling != nil {
		countEl = CountElements(countEl, n.NextSibling)
	}
	return countEl
}

// TextFromHTML get text from html
func TextFromHTML(res []string, n *html.Node) []string {
	if n.Type == html.TextNode {
		res = append(res, n.Data)
	}
	if c := n.FirstChild; c != nil {
		res = TextFromHTML(res, c)
	}
	if n.NextSibling != nil {
		res = TextFromHTML(res, n.NextSibling)
	}

	return res
}

// GetURL try parse html body from URI string
func GetURL(url string) (doc *html.Node, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	doc, err = html.Parse(resp.Body)
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	return doc, nil
}

// CountWordsAndImages count words and images
func CountWordsAndImages(n *html.Node) (words, images int) {
	for _, rawLine := range TextFromHTML(nil, n) {
		line := strings.TrimSpace(rawLine)
		if line == "" {
			continue
		}
		// fmt.Println(line)
		words += len(strings.Fields(line))
	}
	mapElementCount := CountElements(map[string]int{}, n)
	images = mapElementCount["img"]
	return
}
