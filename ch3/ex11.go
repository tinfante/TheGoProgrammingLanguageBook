package main

import (
	"bytes"
	"fmt"
    "strings"
)

func comma(s string) string {
    var decimals string
    decimal_idx := strings.Index(s, ".")
    if decimal_idx != -1 {
        decimals = s[decimal_idx:]
        s = s[:decimal_idx]
    }
    var sign string
    if strings.HasPrefix(s, "-") || strings.HasPrefix(s, "+") {
        sign = string(s[0])
        s =  s[1:]
    }
	if len(s) <= 3 {
		return s
	}
	var buf bytes.Buffer
	rem := len(s) % 3
	if rem != 0 {
		buf.WriteString(s[:rem] + ",")
		s = s[rem:]
	}
	for i, c := range s {
		if i != 0 && i%3 == 0 {
			buf.WriteString(",")
		}
		buf.WriteRune(c)
	}
	return sign + buf.String() + decimals
}

func main() {
	fmt.Println(comma("+123456.33"))
}
