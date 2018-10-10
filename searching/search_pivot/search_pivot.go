package main

import (
	"fmt"
)

func main() {
	arr := []int{5, 6, 7, 8, 9, 10, 1, 2, 3}
	key := 3
	pivotedBinarySearch(arr, len(arr), key)
}

func pivotedBinarySearch(arr []int, length int, key int) int {
	pivot := findPivot(arr, 0, len(arr)-1)
	if pivot == -1 {
		return binarySearch(arr, key)
	}
	if arr[pivot] > arr[0] {
		return binarySearch(arr, key)
	}
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

func findPivot(arr []int, low int, high int) int {
	if high < low {
		return -1
	}
	if high == low {
		return low
	}
	middle := (low + high) / 2
	if middle < high && arr[middle] > arr[middle+1] {
		return middle
	}
	if middle > low && arr[middle] < arr[middle-1] {
		return middle - 1
	}
	if arr[low] >= arr[middle] {
		return findPivot(arr, low, middle-1)
	}
	return findPivot(arr, middle+1, high)
}
