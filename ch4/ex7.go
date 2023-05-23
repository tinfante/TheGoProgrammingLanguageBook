package main

import (
    "fmt"
    "unicode/utf8"
)


func main() {
    s := "hello 世界 world!"
    n := string(reverse([]byte(s)))
    fmt.Println(s)
    fmt.Println(n)
}

func reverse(b []byte) []byte {
    out := make([]byte, len(b))
    i := len(b)
    for i>0 {
        r, s := utf8.DecodeLastRune(b[:i])
        out = append(out, string(r)...)
        i -= s
    }
    return out
}

