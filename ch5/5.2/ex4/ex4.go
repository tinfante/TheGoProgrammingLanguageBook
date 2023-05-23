package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
                links = append(links, "LINK: "+a.Val)
			}
		}
	} else if n.Type == html.ElementNode && n.Data == "script" {
		for _, s := range n.Attr {
			if s.Key == "src" {
                links = append(links, "SCRIPT: "+s.Val)
			}
		}
	} else if n.Type == html.ElementNode && n.Data == "img" {
		for _, i := range n.Attr {
			if i.Key == "src" {
                links = append(links, "IMAGE: "+i.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
