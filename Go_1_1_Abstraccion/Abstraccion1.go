package main

import "fmt"

type Saludo struct {} // Go no tiene el concepto de clases

func (s Saludo) saludar()  {
	fmt.Println("Método en ejecución")
}

func main()  {
	sal := Saludo{}
	sal.saludar()
}