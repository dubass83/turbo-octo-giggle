// Package links provide a funktion to extract links from html doc.
package links

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

// Extract run HTTP-request GET at the required URL
// return links from the HTML document
func Extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("get %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("analyse %s: as HTML: %v", url, err)
	}
	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore incorect links
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

// forEachNode call function pre(x) and post(x) for every x in the tree
// with root n. Both functions are optional
func forEachNode(doc *html.Node, pre, post func(doc *html.Node)) {
	if pre != nil {
		pre(doc)
	}

	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(doc)
	}
}
