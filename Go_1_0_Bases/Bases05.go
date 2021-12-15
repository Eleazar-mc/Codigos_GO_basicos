package main

import "fmt"

func main()  {
	var grado = "B"
	var calificacion = 90

	switch calificacion {
	case 90: grado = "A"
	case 80: grado = "B"
	case 50,60,70 : grado = "C"
	default: grado = "D"
	}
	switch {
	case grado == "A" :
		fmt.Printf("Excelente\n" )
	case grado == "B", grado == "C" :
		fmt.Printf("Bien hecho\n" )
	case grado == "D" :
		fmt.Printf("Pasaste\n" )
	case grado == "F":
		fmt.Printf("Intenta de nuevo\n" )
	default:
		fmt.Printf("Grado no válido\n" )
	}
	fmt.Printf("Tu grado es %s\n", grado)

	var x interface{}

	switch i := x.(type) {
	case nil:
		fmt.Printf("Tipo de x: %T", i)
	case int:
		fmt.Printf("x es int")
	case float64:
		fmt.Printf("x es float64")
	case func(int) float64:
		fmt.Printf("x es func(int)")
	case bool, string:
		fmt.Printf("x es bool o string")
	default:
		fmt.Printf("No sé el tipo")
	}
}