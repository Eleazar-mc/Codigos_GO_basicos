package main

import "fmt"

type Entero int // Alias para un int

func (e Entero) doble() int { // Un método se puede definir sobre tipos no estructurados
	return 2 * int(e) // int() trabaja como una función
}

func main()  {
	fmt.Println("Valor del entero: ", Entero(5))
	fmt.Println("Doble del entero: ", Entero(5).doble())
	fmt.Println(int(6))
}