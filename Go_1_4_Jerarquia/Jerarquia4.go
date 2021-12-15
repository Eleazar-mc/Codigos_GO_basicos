package main

import "fmt"

type IAve interface {
	volar()
	nadar()
}

type Perico struct {}
func(p Perico) volar()  {
	fmt.Println("Perico volando")
}
func(p Perico) nadar()  {
	fmt.Println("Un perico no nada")
}

type Pinguino struct {}
func(p Pinguino) volar()  {
	fmt.Println("Un pingüino no vuela")
}
func(p Pinguino) nadar()  {
	fmt.Println("Pingüino nadando")
}

func prueba_de_vuelo(ave IAve) {
	ave.volar()
}

func main()  {
	pe := Perico{}
	pi := Pinguino{}
	prueba_de_vuelo(pe)
	prueba_de_vuelo(pi)
}