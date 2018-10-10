package main

import (
	"fmt"
)

func main() {

	x := []int{9, 7, 5, 11, 12, 2, 14, 3, 10, 6}
	quickSort(x, 0, len(x)-1)
	fmt.Println("value after sorting", x)
}

func quickSort(arr []int, low int, high int) []int {
	if low < high {
		split := partition(arr, low, high)
		quickSort(arr, low, split-1)
		quickSort(arr, split+1, len(arr)-1)
	}
	return arr
}

func partition(arr []int, low int, high int) int {
	i := low - 1
	pivot := arr[high]

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[high], arr[i+1] = arr[i+1], arr[high]
	return i + 1
}
