package main

import "fmt"

func main() {
    i := 0
    for x := 1; x < 1001; x++ {
        if (x % 3 == 0) || (x % 5 == 0) {
            i = i + x
        }
    }
    fmt.Println(i)
}