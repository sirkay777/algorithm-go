package quicksort

import (
	"testing"
)

func TestQuickSort(t *testing.T) {
	data := []struct {
		input    []int
		expected []int
	}{
		{
			[]int{},
			[]int{},
		},
		{
			[]int{1},
			[]int{1},
		},
		{
			[]int{1, 2},
			[]int{1, 2},
		},
		{
			[]int{2, 1},
			[]int{1, 2},
		},
		{
			[]int{1, 3, 2, 4, 5},
			[]int{1, 2, 3, 4, 5},
		},
	}
	for _, d := range data {
		got := qsort(d.input)
		if !arrSameOrder(got, d.expected) {
			t.Errorf("expected: %v, got: %v\n", d.expected, got)
		}
	}
}

func arrSameOrder(arr1 []int, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}
