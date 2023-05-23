package main

import (
	"fmt"
	"os"
	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
    var elemFreq = make(map[string]int)
    visit(elemFreq, doc)

    for e, f := range elemFreq {
        fmt.Printf("%-10s\t%d\n", e, f)
    }
}

func visit(elemFreq map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
        elemFreq[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(elemFreq, c)
	}
	return elemFreq
}
