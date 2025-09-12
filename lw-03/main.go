package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func worker(id int, dataChan <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for data := range dataChan {
		fmt.Printf("Воркер %d получил данные: %d\n", id, data)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	numWorkers := 1
	if len(os.Args) > 1 {
		n, err := strconv.Atoi(os.Args[1])
		if err == nil && n > 0 {
			numWorkers = n
		}
	}
	fmt.Printf("Запускаем %d воркеров...\n", numWorkers)
	dataChan := make(chan int)
	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, dataChan, &wg)
	}

	for i := 0; ; i++ {
		dataChan <- i
		time.Sleep(100 * time.Millisecond)
	}
}
