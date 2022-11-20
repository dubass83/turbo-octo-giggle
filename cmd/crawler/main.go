package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/dubass83/turbo-octo-giggle/pkg/crawler"
	"golang.org/x/net/html"
)

func main() {
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

	body, err := crawler.GetURL("https://dubass83.xyz")

	if err != nil {
		fmt.Fprintf(os.Stderr, "CountWordsAndImages: %v\n", err)
		os.Exit(1)
	}

	words, images := crawler.CountWordsAndImages(body)
	fmt.Printf("Current link has %d words\n and %d images\n", words, images)
}
