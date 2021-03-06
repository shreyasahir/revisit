package main

import (
	"reflect"
	"testing"
)

var (
	tables = []struct {
		x []int
		y []int
	}{
		{[]int{64, 34, 25, 12, 22, 11, 90}, []int{11, 12, 22, 25, 34, 64, 90}},
	}
)

func TestInsertionSort(t *testing.T) {
	for _, table := range tables {
		sorted := quickSort(table.x, 0, len(table.x)-1)
		if !reflect.DeepEqual(sorted, table.y) {
			t.Errorf("Sorting was incorrect got %d expected %d", sorted, table.y)
		}
	}
}
