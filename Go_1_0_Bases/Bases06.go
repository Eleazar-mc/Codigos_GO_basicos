package main

import "fmt"

func main() {
	var b = 15
	var a int
	numeros := [6]int{1, 2, 3, 5}

	for a := 0; a < 10; a++ {
		fmt.Printf("Valor de a: %d\n", a)
	}
	for a < b {
		a++
		fmt.Printf("Valor de a: %d\n", a)
	}
	for i, x:= range numeros {
		fmt.Printf("Valor de x = %d en índice %d\n", x, i)
	}
}

/* Investigar cómo funcionan:
   break, continue y goto
 */