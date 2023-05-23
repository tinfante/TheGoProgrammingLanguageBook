package main

import (
    "fmt"
    "unicode"
    "unicode/utf8"
)

func squash(strSlice []byte) []byte {
    lastIsSpace := false
    ri, wi := 0, 0
    for ri < len(strSlice) {
        r, s := utf8.DecodeRune(strSlice[ri:])
        if unicode.IsSpace(r) {
            if !lastIsSpace {
                strSlice[wi] = ' '
                wi++
            }
            lastIsSpace = true
        } else {
            utf8.EncodeRune(strSlice[wi:], r)
            wi += s
        }
        ri += s
    }
    return strSlice[:wi]
}

func main() {
    strSlice := "asd \t \n qwe"
    fmt.Println(strSlice, []byte(strSlice))
    squashed := squash([]byte(strSlice))
    fmt.Println(string(squashed), squashed)
}
