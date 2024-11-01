package main

import "fmt"

func inputForBS() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := 6
	fmt.Println(binarysearch(array, target))
}

func binarysearch(array []int, target int) int {
	start := 0
	end := len(array) - 1

	for start <= end {
		mid := start + (end-start)/2

		if target > array[mid] {
			start = mid + 1
		} else if target < array[mid] {
			end = mid - 1
		} else {
			return mid
		}
	}
	return -1
}
