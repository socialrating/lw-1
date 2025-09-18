package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(ctx context.Context, wg *sync.WaitGroup, id int) {
	defer wg.Done()
	fmt.Printf("Воркер %d запущен\n", id)

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Воркер %d останавливается...\n", id)
			return
		default:
			fmt.Printf("Воркер %d работает...\n", id)
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	fmt.Println("Программа запущена. Нажмите Ctrl+C для завершения.")

	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(ctx, &wg, i)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan
	fmt.Println("\nПолучен сигнал прерывания. Инициирую завершение...")

	cancel()

	wg.Wait()

	fmt.Println("Все воркеры завершили работу. Программа корректно остановлена.")
}
