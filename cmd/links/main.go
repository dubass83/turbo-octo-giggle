package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"

	"github.com/dubass83/turbo-octo-giggle/pkg/links"
)

var siteURL string

func init() {
	flag.StringVar(&siteURL, "url", "", "URL string")
}

func main() {
	flag.Parse()
	if siteURL == "" {
		fmt.Fprintln(os.Stderr, "no URL to parse was provided")
		flag.Usage()
		os.Exit(1)
	}
	_, err := url.ParseRequestURI(siteURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "bad URL: %s was provided.\n error: %v", siteURL, err)
	}
	links.Extract(siteURL)
}
