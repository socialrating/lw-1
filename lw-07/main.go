package main

import (
	"fmt"
	"sync"
)

type ConcurrentMap struct {
	mu   sync.Mutex
	data map[string]int
}

func NewConcurrentMap() *ConcurrentMap {
	return &ConcurrentMap{
		data: make(map[string]int),
	}
}

func (cm *ConcurrentMap) Set(key string, value int) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	cm.data[key] = value
}

func (cm *ConcurrentMap) Get(key string) (int, bool) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	val, ok := cm.data[key]
	return val, ok
}

func main() {
	cMap := NewConcurrentMap()

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			key := fmt.Sprintf("key-%d", n)
			cMap.Set(key, n)
		}(i)
	}

	wg.Wait()

	if val, ok := cMap.Get("key-50"); ok {
		fmt.Printf("Значение для key-50: %d\n", val)
	}
	if val, ok := cMap.Get("key-99"); ok {
		fmt.Printf("Значение для key-99: %d\n", val)
	}
	fmt.Printf("Итоговый размер map: %d\n", len(cMap.data))
}
