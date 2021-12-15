package main

import "fmt"

func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main(){
	/* nextNumber es una función con i = 0 */
	nextNumber := getSequence()

	/* Invocación a nextNumber para incrementar i en 1 y devolverla */
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	/* Creación  de una nueva secuencia, i es 0 de nuevo */
	nextNumber1 := getSequence()
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())
}

// Este ejemplo representa los clousures, funciones que recuerdan el contexto de ejecución