package main

import (
	"fmt"
)

func intersection(a, b []int) []int {
	m := make(map[int]bool)
	result := []int{}
	seen := make(map[int]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if m[item] && !seen[item] {
			result = append(result, item)
			seen[item] = true
		}
	}
	return result
}

func main() {
	A := []int{1, 2, 3}
	B := []int{2, 3, 4}

	intersect := intersection(A, B)
	fmt.Println(intersect)
}
