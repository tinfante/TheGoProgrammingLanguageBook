package main

import (
    "fmt"
    "io"
    "os"
    "strings"
    "golang.org/x/net/html"
)


func main() {
    stdin, err := io.ReadAll(os.Stdin)
    if err != nil {
        panic(err)
    }
    in := string(stdin)
    _, err2 := html.Parse(strings.NewReader(in))
	if err2 != nil {
		panic(err)
	}
    fmt.Println("HTML from Stdin parsed correctly.")
}
