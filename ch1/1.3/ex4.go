package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
)


func main() {
    counts := make(map[string]int)
    found := make(map[string][]string)
    files := os.Args[1:]
    if len(files) == 0 {
        countLines(os.Stdin, counts, found, false)
    } else {
        for _, arg := range files {
            f, err := os.Open(arg)
            if err != nil {
                fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
                continue
            }
            countLines(f, counts, found, true)
            f.Close()
        }
    }
    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\t%v\n", n, line, strings.Join(found[line], ", "))
        }
    }
}

func countLines(f *os.File, counts map[string]int, found map[string][]string, isFile bool) {
    input := bufio.NewScanner(f)
    for input.Scan() {
        line := input.Text()
        counts[line]++
        var fn string
        if isFile == true {
            fn = f.Name()
        } else {
            fn = "stdin"
        }
        if isInSlice(found[line], fn) {
            continue
        }
        found[line] = append(found[line], fn)
    }
}

func isInSlice(sl []string, v string) bool {
    op := false
    for _, e := range sl {
        if v == e {
            op = true
            break
        }
    }
    return op
}
