package main

import "fmt"

func factorial() {
	var k int
	fact := 1
	fmt.Println("Enter the number:")
	fmt.Scanln(&k)
	fmt.Println("hello")
	for i := k; i > 0; i-- {
		fact = fact * i
	}
	fmt.Println("The factorial of: ", k, "is:", fact)
}
