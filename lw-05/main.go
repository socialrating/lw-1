package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	const N = 2

	go func() {
		i := 0
		for {
			ch <- i
			i++
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		for {
			value := <-ch
			fmt.Printf("Получено значение: %d\n", value)
		}
	}()

	fmt.Printf("Программа завершится через %d секунд...\n", N)
	<-time.After(N * time.Second)

	fmt.Println("Время вышло. Программа завершена.")
}
