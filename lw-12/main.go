package main

import (
	"fmt"
)

func main() {
	data := []string{"cat", "cat", "dog", "cat", "tree"}

	set := make(map[string]struct{})
	for _, item := range data {
		set[item] = struct{}{}
	}

	result := make([]string, 0, len(set))
	for key := range set {
		result = append(result, key)
	}

	fmt.Println(result)
}
