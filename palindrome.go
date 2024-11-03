package main

func palindromeCheck(number int) bool {
	original := number

	reverse := 0

	for number > 0 {
		digits := number % 10
		reverse = reverse*10 + digits
		number /= 10
	}
	return reverse == original
}
