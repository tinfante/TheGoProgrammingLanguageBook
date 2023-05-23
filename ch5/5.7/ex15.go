package main

import (
    "fmt"
    "os"
)


func max(values ...int) int {
    if len(values) == 0 {
        os.Exit(1)
    }
    largest := values[0]
    for i, v := range values {
        if i == 0 {
            continue
        }
        if  v > largest {
            largest = v
        }
    }
    return largest
}

func min(values ...int) int {
    if len(values) == 0 {
        os.Exit(1)
    }
    smallest := values[0]
    for i, v := range values {
        if i == 0 {
            continue
        }
        if  v < smallest {
            smallest = v
        }
    }
    return smallest
}

func main() {
    values :=  []int{4, 6, 1, 7, 3}
    fmt.Println("max", values, "is", max(values...))
    fmt.Println("min", values, "is", min(values...))
}
