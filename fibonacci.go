package main

import "fmt"

func fibo(k int) int {
	if k < 0 {
		return 0
	} else if k == 1 {
		return 1
	}
	return fibo(k-1) + fibo(k-2)

}

func fiborange(k int) {

	if k < 0 {
		fmt.Println("Please enter a poistive number!!")
	}
	a, b := 0, 1
	fmt.Println("Printing fibonacci series")
	for i := 0; i < k; i++ {
		if i == 0 {
			fmt.Println(a, "")
			continue
		}
		if i == 1 {
			fmt.Println(b, "")
			continue
		}
		next := a + b
		fmt.Println(next, "")
		a, b = b, next
	}
	fmt.Println()
}

//   0 1  2 3 5 8 13g
