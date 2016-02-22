package main

import (
        "fmt"
)

func oneTwoThree() (int, int, int) {
        return 1, 2, 3
}

func main() {
        a, _, c := oneTwoThree()
        fmt.Printf("%d %d\n", a, c)
}