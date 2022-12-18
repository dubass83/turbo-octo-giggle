package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/dubass83/turbo-octo-giggle/pkg/crawler"
	"golang.org/x/net/html"
)

var (
	textFromHTML        bool
	countByEllement     bool
	countWordsAndImages bool
	forEachNode         bool
	fetch               bool
)

func init() {
	flag.BoolVar(&textFromHTML, "textFromHTML", false, "get text from html")
	flag.BoolVar(&countByEllement, "countByEllement", false, "count by element")
	flag.BoolVar(&countWordsAndImages, "countWordsAndImages", false, "count words and images")
	flag.BoolVar(&forEachNode, "forEachNode", false, "print each node")
	flag.BoolVar(&fetch, "fetch", true, "save HTTP response to the file")
}

func main() {
	// Call flag.Parse() to parse command-line flags
	flag.Parse()

	if textFromHTML {
		doc, err := html.Parse(os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
			os.Exit(1)
		}
		for _, rawLine := range crawler.TextFromHTML(nil, doc) {
			line := strings.TrimSpace(rawLine)
			if line == "" {
				continue
			}
			fmt.Println(line)
		}
	}

	if countByEllement {
		doc, err := html.Parse(os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
			os.Exit(1)
		}
		mapElementCount := crawler.CountElements(map[string]int{}, doc)
		sliceKeysMap := make([]string, 0, len(mapElementCount))
		for key := range mapElementCount {
			sliceKeysMap = append(sliceKeysMap, key)
		}
		sort.Slice(sliceKeysMap, func(x, y int) bool {
			return mapElementCount[sliceKeysMap[x]] > mapElementCount[sliceKeysMap[y]]
		})
		for _, val := range sliceKeysMap {
			fmt.Printf("[%-7s]=>%d\n", val, mapElementCount[val])
		}
	}

	if countWordsAndImages {
		body, err := crawler.GetURL("https://dubass83.xyz")

		if err != nil {
			fmt.Fprintf(os.Stderr, "CountWordsAndImages: %v\n", err)
			os.Exit(1)
		}

		words, images := crawler.CountWordsAndImages(body)
		fmt.Printf("Current link has %d words\n and %d images\n", words, images)
	}

	if forEachNode {
		doc, err := html.Parse(os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
			os.Exit(1)
		}
		crawler.ForEachNode(doc, crawler.StartElement, crawler.EndElement)
	}

	if fetch {
		url := "https://dubass83.xyz/"
		filename, n, err := crawler.Fetch(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v", err)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stdout, "save URL: %s to file %s\nbyte: %d", url, filename, n)
	}
}
