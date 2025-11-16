package main

import "fmt"

var justString string

func someFunc() {
    v := createHugeString(1 << 10)

    temp := make([]byte, 100)
    copy(temp, v[:100])
    justString = string(temp)
}

func createHugeString(size int) string {
    return string(make([]byte, size))
}

func main() {
    someFunc()
    fmt.Println(len(justString))
}
