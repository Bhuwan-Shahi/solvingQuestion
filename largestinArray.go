package main

import (
	"fmt"
)

func inputforLargest() {
	array := []int{2, 4, 12, 1, 19, 8, 56, 100}
	fmt.Println("This the largest in array:")
	fmt.Println(largest(array))
	fmt.Println("This the smallest in array:")
	fmt.Println(smallest(array))

}
func smallest(array []int) int {
	Smallest := array[0]
	for _, value := range array {
		if value < Smallest {
			Smallest = value
		}
	}
	return Smallest

}

func largest(array []int) int {
	Largest := array[0]
	for _, value := range array {
		if value > Largest {
			Largest = value
		}
	}
	return Largest

}
