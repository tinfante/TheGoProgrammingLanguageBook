package main

import "fmt"

// remove adjacent duplicates in-place.
func rad(strSlice []string) []string {
    var lastStr string
    i := 0
    for _, s := range strSlice {
        if s != lastStr {
            strSlice[i] = s
            lastStr = s
            i++
        }
    }
    return strSlice[:i]
}

func main() {
    strSlice := []string{"a", "b", "c", "c", "d", "e", "f", "f", "f",  "g", "a", "b", "c"}
    fmt.Println(strSlice)
    strSlice = rad(strSlice)
    fmt.Println(strSlice)
}
