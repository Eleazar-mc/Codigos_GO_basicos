package main

import "fmt"

type Base struct {
	colorbase string
}

func(s *Base) m()  {
	fmt.Println("Base.m()")
}

type Derivada struct {
	Base
	colorderivada string
}

func imposible(b Base) { // Cambiar por tipo IBase
	b.m()
}

func main() {
	b := Base{colorbase:"rojo"}
	d := &Derivada{ // Asignar por la dirección: &Derivada
		Base:          b,
		colorderivada: "verde",
	}
	d.m()
	fmt.Println("El color en Base es:", d.colorbase)
	fmt.Println("Y el color en Derivada es:", d.colorderivada)
	imposible(b) // Asignar la dirección: &b
	//imposible(d) // Error, no hay jerarquía
}

type IBase interface {
	m()
}