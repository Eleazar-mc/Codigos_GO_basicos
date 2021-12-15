package main

import "fmt"

func main() {
	var x = [5] int {100,200,300,400,500} // Definición de una rebanada (slice)
	fmt.Println("Longitud de x: ", len(x))
	fmt.Println("Capacidad de x: ", cap(x)) // Capacidad de acomodar elementos
	fmt.Println("Valores de x: ", x)

	var y [] int // Rebanada sin tamaño especificado
	fmt.Println("Longitud de y: ", len(y))
	fmt.Println("Capacidad de y:", cap(y))
	fmt.Println("Valores de y: ", y)

	var z = make([]int, 3, 5)
	fmt.Println("Longitud de z: ", len(z))
	fmt.Println("Capacidad de z: ", cap(z))
	fmt.Println("Valores de z: ", z)
	z[1] = 9
	fmt.Println(z[1:3]) // Índices 1 inclusivo, 3 exclusivo
	// ¿Cómo acceder a la capacidad máxima?

	numeros := []int{0,10,20,30,40,50,6,7,8}
	for i:= range numeros {
		fmt.Println("Elemento del slice",i,"es", numeros[i])
	}

	// ¿Los índices del arreglo y del slice son iguales?
}