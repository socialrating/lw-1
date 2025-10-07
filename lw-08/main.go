package main

import "fmt"

func setBitToOne(n int64, i int) int64 {
    return n | (1 << i)
}

func setBitToZero(n int64, i int) int64 {
    return n &^ (1 << i)
}

func main() {
    num := int64(5)
    i := 1

    fmt.Printf("Original number: %d\n", num)

    newNumZero := setBitToZero(num, i)
    fmt.Printf("After setting %d-th bit to 0: %d\n", i, newNumZero)

    newNumOne := setBitToOne(num, i)
    fmt.Printf("After setting %d-th bit to 1: %d\n", i, newNumOne)
}
