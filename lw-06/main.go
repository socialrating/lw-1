package main

import (
	"fmt"
	"time"
)

func workerNaturalExit(tasks <-chan int) {
	fmt.Println("Воркер запущен, ждет задач...")
	for task := range tasks {
		fmt.Printf("Получена задача: %d\n", task)
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println("Все задачи выполнены. Воркер остановлен.")
}

func main() {
	tasks := make(chan int, 3)

	go workerNaturalExit(tasks)

	for i := 1; i <= 3; i++ {
		fmt.Printf("Отправлена задача %d\n", i)
		tasks <- i
	}

	close(tasks)

	time.Sleep(2 * time.Second)
	fmt.Println("Программа завершена.")
}
