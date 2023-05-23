package main

import (
    "fmt"
    "net/http"
    "strings"
    "strconv"
    "os"
    "golang.org/x/net/html"
)


func CountWordsAndImages(url string) (words, images int, err error) {
    resp, err := http.Get(url)
    if err != nil {
        return
    }
    doc, err := html.Parse(resp.Body)
    resp.Body.Close()
    if err != nil {
        err = fmt.Errorf("parsing HTML: %s", err)
        return
    }
    words, images = countWordsAndImages(doc, 0, 0)
    return words, images, nil
}


func  countWordsAndImages(n *html.Node, w, i int) (words, images int) {
    if n.Type == html.ElementNode && n.Data == "img" {
        i++
    } else if n.Type == html.TextNode && strings.TrimSpace(n.Data) != "" {
        for range strings.Split(strings.TrimSpace(n.Data), " ") {
            w++
        }
    }
    for c := n.FirstChild; c != nil; c = c.NextSibling {
        w, i = countWordsAndImages(c, w, i)
    }
    return w, i
}


func main() {
    words, images, _ := CountWordsAndImages(os.Args[1])
    fmt.Println("words: " + strconv.Itoa(words) + "\nimages: " + strconv.Itoa(images))
}

