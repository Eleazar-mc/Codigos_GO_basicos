package main

import "fmt"

func max(num1, num2 int) int { // El tipo de retorno se especifica después de los paréntesis
	var resultado int

	if num1 > num2 {
		resultado = num1
	} else {
		resultado = num2
	}
	return resultado
}

func main()  {
	var a = 100
	var b = 200
	var ret int

	ret = max(a, b)
	fmt.Printf( "El valor máximo es : %d\n", ret)
}