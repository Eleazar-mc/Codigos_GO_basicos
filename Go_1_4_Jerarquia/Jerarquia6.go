package main

import "fmt"

type Volador struct {}
func(v Volador) volar()  {
	fmt.Println("Volando")
}

type Extraterrestre struct {}
func(e Extraterrestre) golpear()  {
	fmt.Println("Golpenado")
}

type Debil struct {}
func(d Debil) debilitar()  {
	fmt.Println("¡La fuerza se reduce!")
}

type Superman struct {
	Volador
	Extraterrestre
	Debil
}
func(s Superman) defender()  {
	fmt.Println("Defendiendo el planeta")
}

func main()  {
	s := Superman{Volador{}, Extraterrestre{}, Debil{}}
	s.volar()
	s.defender()
	s.golpear()
	s.debilitar()
}