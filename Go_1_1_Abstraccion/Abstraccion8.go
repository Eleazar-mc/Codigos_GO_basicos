package main

import "fmt"

type Abstracto struct {
	nombre string
}

type Concreto struct {
	Abstracto
}

func (a *Abstracto) concreto1() {
	fmt.Println("Abstracto.concreto1()")
}

func (b *Concreto) concreto2() {
	fmt.Println("Concreto.concreto2(), y nombre", b)
}

func main() {
	a := Abstracto{}
	b := &Concreto{
		Abstracto: a,
	}
	b.nombre = "Go"
	b.concreto1()
	b.concreto2()
	fmt.Println(b.nombre)
}