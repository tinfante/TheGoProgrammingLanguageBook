package main

import (
    "fmt"
    "bufio"
    "os"
)

func main() {
    wordFreq := make(map[string]int)
    in := bufio.NewScanner(os.Stdin)
    in.Split(bufio.ScanWords)
    for in.Scan() {
        wordFreq[in.Text()]++
    }
    fmt.Println("word\tfreq")
    for w, f := range wordFreq {
        fmt.Printf("%s\t%d\n", w, f)
    }
}
