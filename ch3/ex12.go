package main

import (
    "fmt"
    "bytes"
)

func isAnagram(s string) bool {
    left := s[:len(s)/2]
    var right string
    mid := len(s) % 2
    if mid == 0 {
        right = s[len(s)/2:]
    } else {
        right = s[(len(s)/2)+1:]
    }
    right = reverse(right)
    if left == right {
        return true
    }
    return false
}


func reverse(s string) string {
	var buf bytes.Buffer
    for i:=len(s)-1; i>=0; i-- {
        buf.WriteByte(s[i])
    }
    return buf.String()
}

func main() {
    s := "qweewq"
    fmt.Println(s, isAnagram(s))
}
