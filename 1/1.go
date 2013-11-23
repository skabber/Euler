package main

import "fmt"

func main() {
    i := 0
    sum := 0
    for i < 1001 {
        i++
        if (i % 3 == 0) || (i % 5 == 0) {
            sum = sum + i
        }

    }
    fmt.Println(sum)
}