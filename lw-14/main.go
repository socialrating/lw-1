package main

import "fmt"

func detectType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan int:
		fmt.Println("chan int")
	default:
		fmt.Println("unknown")
	}
}

func main() {
	detectType(42)
	detectType("hello")
	detectType(true)
	ch := make(chan int)
	detectType(ch)
	detectType(3.14)
}
