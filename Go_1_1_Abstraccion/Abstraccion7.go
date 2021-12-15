package main

import "fmt"

type Puntos struct {
	x, y float64
}

type Punto3Ds struct {
	Puntos // Tipo embebido
	z float64
}

func main()  {
	p := Puntos{x: 0, y: 0}
	fmt.Println(p)
	p.x	= -0.258
	p.y = 23.158
	fmt.Println("Punto con coordenadas: ", p)

	q := Punto3Ds{
		Puntos: Puntos{x:0, y:0},
		z:      0,
	}
	fmt.Println(q)
	q.x = 5.698
	q.y = -0.002
	q.z = 1.578
	fmt.Println("Punto 3D con coordenadas: ", q)
	fmt.Println("Punto 3D con coordenadas:", q.x, q.y, q.z)
	fmt.Println(Punto3Ds.String(q))
}

func (p Punto3Ds) String() string {
	return fmt.Sprint("(", p.x, ",", p.y, "):", p.z)
}