package main

import (
	"fmt"
)

func main() {
	arr := []int{10, 12, 13, 16, 18, 19, 20, 21, 22, 23, 24, 33, 35, 42, 47}
	fmt.Println(interpolationSearch(arr, 16))
}

func interpolationSearch(arr []int, x int) int {
	low := 0
	high := len(arr) - 1
	for low <= high && x > arr[low] && x < arr[high] {
		pos := low + (((high - low) / (x - arr[low])) * (arr[high] - arr[low]))

		// int pos = lo + (((hi-lo) /
		// (arr[hi]-arr[lo]))*(x - arr[lo]));
		fmt.Println("pos", pos)
		if arr[pos] == x {
			return pos
		}
		if arr[pos] < x {
			low = pos + 1
		} else {
			high = pos - 1
		}
	}

	return -1
}
