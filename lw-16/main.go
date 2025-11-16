package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[len(arr)/2]

	var left []int
	var right []int
	var equal []int

	for _, v := range arr {
		if v < pivot {
			left = append(left, v)
		} else if v > pivot {
			right = append(right, v)
		} else {
			equal = append(equal, v)
		}
	}

	left = quickSort(left)
	right = quickSort(right)

	return append(append(left, equal...), right...)
}

func main() {
	arr := []int{5, 2, 9, 1, 5, 6}

	fmt.Println("Исходный:", arr)
	sorted := quickSort(arr)
	fmt.Println("Отсортированный:", sorted)
}
