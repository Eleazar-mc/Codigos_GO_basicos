package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var x = rand.Float64()
	fmt.Println(x)
	if(x < 0.5) {
		fmt.Println("x es menor a 0.5")
	} else { // else es obligatorio que esté en la misma línea de cierre del if
		fmt.Println("x es mayor o igual a 0.5")
	}
}