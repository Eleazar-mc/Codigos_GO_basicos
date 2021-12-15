package main

import (
	"fmt"
	"math"
)

type Circulo struct { // Nuevo tipo de dato
	x, y, radio float64
}

func(circle Circulo) area() float64 { // Nuevo método para el círculo
	return math.Pi * math.Pow(circle.radio, 2)
}

func main() {
	c := Circulo{x: 0, y: 0, radio: 5}
	fmt.Printf("Área del círculo: %f\n", c.area())

	d := Circulo{y: 0, radio: 5}
	fmt.Printf("Área del círculo: %f\n", d.area())

	e := Circulo{}
	fmt.Printf("Área del círculo: %f\n", e.area())
}