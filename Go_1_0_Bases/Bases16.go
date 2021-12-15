package main

import ("fmt"; "math")

type Shape interface { // Definición de una interfaz
	area() float64
}

type Circunferencia struct { // Definición de un círculo
	x, y, radio float64
}

type Rectangulo struct { // Definición de un rectángulo
	ancho, alto float64
}

func(cir Circunferencia) area() float64 { // Método: implementación de Shape.area()
	return math.Pi * cir.radio * cir.radio
}

func(rect Rectangulo) area() float64 { // Método: implementación de Shape.area()
	return rect.ancho * rect.alto
}

func getArea(shape Shape) float64 { // Definición de un método para shape
	return shape.area()
}

func main() {
	circulo := Circunferencia{x:0, y:0, radio:5}
	rectangulo := Rectangulo{ancho: 10, alto:5}

	fmt.Printf("Área del círculo: %f\n", getArea(circulo))
	fmt.Printf("Área del rectángulo: %f\n", getArea(rectangulo))
}