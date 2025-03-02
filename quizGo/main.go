package main 

import (
	"fmt"
)

func sortDescending(slice *[]int) {
	arr := *slice

	n := len(arr)
	for i := 0; i < n - 1; i++ {
		for j := 0; j < n - i - 1; j++ {
			if arr[j] < arr[j + 1] {
				arr[j], arr[j + 1] = arr[j + 1], arr[j]
			}
		}
	}
}

func main() {
	arr := []int{12, 4, 5, 3, 25, 794, 867, 1, 74, 19, 16, 82, 100, 391}
	fmt.Printf("\nArray before sorting: %v\n\n", arr)

	sortDescending(&arr)
	fmt.Printf("        Sorted array: %v\n", arr)
}