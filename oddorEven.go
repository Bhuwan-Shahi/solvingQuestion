package main

import "fmt"

func oddorEven() {
	var k int
	fmt.Println("Enter you num:")
	fmt.Scanln(&k)
	fmt.Println(k)

	if k%2 == 0 {
		fmt.Println(k, " This is even number!")
	} else {
		fmt.Println("This is odd number")
	}
}
