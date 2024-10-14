package main

import (
	"QuickSort/sorting"
	"fmt"
)

func main() {

	quickSort := sorting.QuickSort{
		Data: []int{3, 7, 8, 5, 2, 9, 6, 4},
	}

	quickSort.Sort(0, len(quickSort.Data)-1)

	fmt.Println(quickSort.Data)
}
