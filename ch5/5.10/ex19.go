package main

import (
    "fmt"
)

func f() (r int){
    defer func() {
        recover()
        r = 9
    }()
    panic("panic")
}

func main() {
    fmt.Println(f())
}
