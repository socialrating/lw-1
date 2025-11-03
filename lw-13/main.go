package main

import "fmt"

func main() {
    a := 5
    b := 7

    fmt.Println("До:", a, b)

    a = a + b 
    b = a - b 
    a = a - b 

    fmt.Println("После:", a, b)
}
