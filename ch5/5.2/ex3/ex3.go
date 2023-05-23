package main

import (
	"fmt"
	"os"
	"golang.org/x/net/html"
    "strings"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
    texts := visit(nil, doc)
    fmt.Println(texts)
    for _, t := range texts {
        fmt.Println(t)
    }

}

func visit(texts []string, n *html.Node) []string {
	if n.Type == html.ElementNode &&
        (n.Data == "style" || n.Data == "script") {
        return texts
	}
    if n.Type == html.TextNode && strings.TrimSpace(n.Data) != "" {
        texts = append(texts, strings.TrimSpace(n.Data))
    }
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		 texts = visit(texts, c)
	}
	return texts
}
