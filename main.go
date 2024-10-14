package main

import (
	"QuickSort/sorting"
	"fmt"
)

func main() {

	quickSort := sorting.QuickSort{
		Data: []int{4, 3, 1, 2, 5, 9, 7, 10, 6},
	}

	quickSort.Sort(4, len(quickSort.Data)-1)

	fmt.Println(quickSort.Data)
}
