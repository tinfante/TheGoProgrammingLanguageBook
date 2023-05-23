package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
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
	return buf.String()
}

func main() {
	fmt.Println(comma("1234567"))
}
