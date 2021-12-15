package main

import ("fmt"; "math")

func main(){
	// Función como variable, función creada al vuelo
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	// Llamada a la función
	fmt.Println(getSquareRoot(9))
}