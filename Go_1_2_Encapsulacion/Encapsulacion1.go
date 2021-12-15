package main

import "fmt"

type sumaXY func(float64, float64) float64

type Punto struct {
	x, y   float64
	operxy sumaXY
}

func main() {
	p := Punto{
		x:      2,
		y:      5,
		operxy: func(x, y float64) float64 {
			return x + y
		},
	}

	fmt.Println(p)
	fmt.Println(p.operxy(p.x, p.y))

	q := Punto{
		x:      -2,
		y:      -5,
		operxy: func(x, y float64) float64 {
			return x / y
		},
	}

	fmt.Println(q)
	fmt.Println(q.operxy(q.x, 19.71))
}