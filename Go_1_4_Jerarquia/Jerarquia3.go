package main

import "fmt"

type IBases interface {
	m()
}

type Bases struct {
	colorbase string
	p func()
}

func(s *Bases) m()  {
	s.p()
}

type Derivadas struct {
	Bases
	colorderivada string
}

func imposibles(b IBases) {
	b.m()
}

func main() {
	b := Bases{
		colorbase:"rojo",
		p: func() {
			fmt.Println("Derivada.p()")
		},
	}
	d := &Derivadas{
		Bases:          b,
		colorderivada: "verde",
	}
	d.m() // Se muestra el método p()
	      // El método p() le pertenece a Base
}