package main

import "fmt"

func factorial(i int) int {
	if i <= 1 {
		return 1
	}
	return i * factorial(i - 1)
}

func main() {
	var i = 33
	fmt.Printf("Factorial de %d es %d", i, factorial(i))
}