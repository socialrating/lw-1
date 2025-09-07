package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup

	for _, num := range numbers {
		wg.Add(1)

		go func(n int) {
			defer wg.Done()

			square := n * n

			fmt.Printf("Квадрат числа %d равен %d\n", n, square)
		}(num)
	}

	wg.Wait()

	fmt.Println("Все вычисления завершены.")
}
