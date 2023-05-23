package main

import (
    "fmt"
)

func join(sep string, values ...string) string {
    joined := values[0]
    for i, v := range values {
        if i == 0 {
            continue
        }
        joined += sep + v
    }
    return joined
}

func main() {
    fmt.Println(join("/", "hola", "mundo", "cruel"))
}
