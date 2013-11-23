package main

import "fmt"

func main() {
    first := 0
    next := first + 1
    sum := 0
    fib := 0

    for {
        fib = first + next

        if (fib > 4000000) {
            fmt.Println(sum)
            break
        }

        if (fib % 2 == 0) {
            sum = sum + fib
        }

        first = next
        next = fib
    }
}