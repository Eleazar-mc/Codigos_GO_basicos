package main

import "fmt"

type IBasex interface {
	m()
}

type Basex struct {
	colorbase string
}

func(s *Basex) m()  {
	s.p()
}

func(s *Basex) p()  {
	fmt.Println("Base.p()")
}

type Derivadax struct {
	Basex
	colorderivada string
}

func imposiblex(b IBasex) {
	b.m()
}

func(s *Derivadax) p()  { // Nuevo método p() para Derivada
	fmt.Println("Derivada.p()")
}

func main() {
	b := Basex{colorbase:"rojo"}
	d := &Derivadax{
		Basex:          b,
		colorderivada: "verde",
	}
	d.m() // Se muestra el método de Base
	d.p() // Derivada.p() debe invocarse explícitamente
}