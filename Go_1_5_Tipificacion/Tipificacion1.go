package main

import "fmt"

type IBase1 interface {
	m()
}

type IBase2 interface {
	p()
}

type Base1 struct {
	color string
}

func(b Base1) m()  {
	fmt.Println("Base1.m()")
}

type Base2 struct {}

func(b *Base2) p()  {
	fmt.Println("Base2.p()")
}

type Derivada struct {
	Base1
	Base2
	color string
}

func(d Derivada) a()  {
	fmt.Println("Derivada.a()")
}

func verifica1(b IBase1)  {
	b.m()
}

func verifica2(b IBase2)  {
	b.p()
}

func main()  {
	base1 := Base1{color:"Rojo"}
	base2 := Base2{}
	deriv := &Derivada{
		Base1: base1,
		Base2: base2,
		color: "Verde",
	}
	deriv.m()
	deriv.p()
	deriv.a()
	verifica1(deriv)
	verifica2(deriv)
}