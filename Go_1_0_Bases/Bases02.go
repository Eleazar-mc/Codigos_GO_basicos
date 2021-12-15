package main

import "fmt"

func main()  {
	var i int
	var f float32 = 2.35
	var b byte
	j := 77 // Declaración corta, inferencia
	var x, y, z = 3, 3.14159, "Go" // Declaración mezclada
	const E = 2.71828

	fmt.Println("Valor de i: ", i)
	fmt.Println(f)
	fmt.Println(b)
	fmt.Println(j)
	fmt.Println(x, y, z)
	fmt.Println(E)
	fmt.Printf("Tipo de f: %T", f)
}