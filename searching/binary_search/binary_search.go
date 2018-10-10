package main

import (
	"fmt"
)

func main() {
	arr := []int{2, 3, 4, 10, 40}
	fmt.Println(binarySearch(arr, 10))
}

func binarySearch(arr []int, elem int) int {
	start := 0
	end := len(arr) - 1
	var middle int
	for start <= end {
		middle = start + (end-start)/2
		if arr[middle] == elem {
			return middle
		} else if arr[middle] < elem {
			start = middle + 1
		} else {
			end = middle - 1
		}
	}
	return -1
}
