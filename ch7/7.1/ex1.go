package main

import (
    "fmt"
    "bufio"
    "strings"
)


type ByteCounter int
type WordCounter int
type LineCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
    *c += ByteCounter(len(p))
    return len(p), nil
}

func (w *WordCounter) Write(p []byte) (int, error) {
    scanner := bufio.NewScanner(strings.NewReader(string(p)))
    scanner.Split(bufio.ScanWords)
    count := 0
    for scanner.Scan() {
       count++
    }
    *w += WordCounter(count)
    return count, nil
}

func (l *LineCounter) Write(p []byte) (int, error) {
    scanner := bufio.NewScanner(strings.NewReader(string(p)))
    scanner.Split(bufio.ScanLines)
    count := 0
    for scanner.Scan() {
       count++
    }
    *l += LineCounter(count)
    return count, nil
}

func main() {
    var c ByteCounter
    c.Write([]byte("This is a sentence."))
    fmt.Println("Byte Counter ", c)

    var w WordCounter
    w.Write([]byte("This is a sentence."))
    fmt.Println("Word Counter ", w)

    var l LineCounter
    l.Write([]byte("This\nis\na\nsentence\n."))
    fmt.Println("Line Counter ", l)
}
