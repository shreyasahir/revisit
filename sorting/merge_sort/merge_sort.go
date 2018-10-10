package main

import (
	"fmt"
)

func main() {
	var x = []int{5, 2, 4, 6, 1, 3}
	result := mergeSort(x)
	fmt.Println(result)
}

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	var left, right []int
	middle := len(arr) / 2
	left = mergeSort(arr[:middle])
	right = mergeSort(arr[middle:])

	return merge(left, right)
}

func merge(left []int, right []int) []int {
	var result []int

	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	if len(left) > 0 {
		result = append(result, left...)
	}

	if len(right) > 0 {
		result = append(result, right...)
	}

	return result
}
