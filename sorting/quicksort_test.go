package sorting

import (
	"reflect"
	"testing"
)

func TestQuickSort_Sort(test *testing.T) {
	sorts := []struct {
		input    []int
		expected []int
	}{
		{input: []int{5, 2, 8, 1, 9}, expected: []int{1, 2, 5, 8, 9}},
		{input: []int{10, 7, 8, 9, 1, 5}, expected: []int{1, 5, 7, 8, 9, 10}},
		{input: []int{3, 7, 8, 5, 2, 1, 9, 6, 4}, expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{input: []int{9, 8, 7, 6, 5, 4, 3, 2, 1}, expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{input: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
	}

	for _, sort := range sorts {

		quickSort := QuickSort{
			Data: sort.input,
		}

		quickSort.Sort(0, len(quickSort.Data)-1)

		if !reflect.DeepEqual(quickSort.Data, sort.expected) {
			test.Errorf("QuickSort(%v); Expected = %v", sort.input, sort.expected)
		}

	}

}
