package main

type IPuedePelear interface {
	pelear()
}
type PuedePelear struct {}
func(p PuedePelear) pelear()  {}

type IPuedeNadar interface {
	nadar()
}
type PuedeNadar struct {}
func(p PuedeNadar) nadar()  {}

type IPuedeVolar interface {
	volar()
}
type PuedeVolar struct {}
func(p PuedeVolar) volar()  {}

type Accion struct {}
func(a Accion) pelear()  {}

type Heroe struct {
	Accion
	PuedeNadar
	PuedePelear
	PuedeVolar
}
func(h Heroe) nadar() {}
func(h Heroe) volar() {}

func t(x IPuedePelear) {
	x.pelear()
}

func u(x IPuedeNadar) {
	x.nadar()
}

func v(x IPuedeVolar) {
	x.volar()
}

func w(x Accion) {
	x.pelear()
}

func main() {
	h := Heroe{
		Accion:      Accion{},
		PuedeNadar:  PuedeNadar{},
		PuedePelear: PuedePelear{},
		PuedeVolar:  PuedeVolar{},
	}
	//t(h)
	u(h)
	v(h)
	//w(h)
}