package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}
func main() {
	a, b := swap("Mahesh", "Kumar")
	fmt.Println(a, b)
}

// Go soporta el paso por valor y paso por referencia para los argumentos