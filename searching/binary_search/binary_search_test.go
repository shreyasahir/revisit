package main

import (
	"testing"
)

var (
	tables = []struct {
		x int
		y int
	}{
		{2, 0},
		{3, 1},
		{10, 3},
		{600, 8},
		{900, -1},
	}
)

func TestBinarySearch(t *testing.T) {
	arr := []int{2, 3, 4, 10, 40, 60, 90, 400, 600}
	for _, table := range tables {
		elem := binarySearch(arr, table.x)
		if elem != table.y {
			t.Errorf("Searching was incorrect got: %d expected: %d", elem, table.y)
		}
	}
}
