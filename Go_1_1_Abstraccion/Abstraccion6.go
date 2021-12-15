package main

import "fmt"

type Punto struct {
	x, y float64
}

type Punto3D struct {
	beast Punto
	z float64
}

func main()  {
	p := Punto{x: 0, y: 0}
	fmt.Println(p)
	p.x	= -0.258
	p.y = 23.158
	fmt.Println("Punto con coordenadas: ", p)

	q := Punto3D{
		beast: Punto{x:99},
		z:     100,
	}
	fmt.Println(q)
	q.beast.x = 5.698
	q.beast.y = -0.002
	q.z = 1.578
	fmt.Println("Punto 3D con coordenadas:", q)
	fmt.Println("Punto 3D con coordenadas:", q.beast.x, q.beast.y, q.z)
	fmt.Println(Punto3D.String(q))
}

func (p Punto3D) String() string {
	return fmt.Sprint("(", p.beast.x, ",", p.beast.y, "):", p.z)
}