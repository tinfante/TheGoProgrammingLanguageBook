package main

import "fmt"

func rotate(s []int, n int) []int {
    var d = make([]int, len(s))
    for i := 0; i < len(s); i++ {
        k := (i-n) % len(s)
        if k < 0 {
            k += len(s)
        }
        d[k] = s[i]
    }
    return d
}

func main() {
    s := []int{1,2,3,4,5,6}
    fmt.Println(s)
    s = rotate(s, 3)
    fmt.Println(s)
}
