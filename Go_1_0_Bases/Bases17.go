package main

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	CeroAbsolutoC Celsius = -273.15
	CongelacionC  Celsius = 0
	EbullicionC   Celsius = 100
)

func CaF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}
func FaC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
func main() {
	fmt.Printf("%g\n", EbullicionC - CongelacionC) // "100" °C
	ebullicionF := CaF(EbullicionC)
	fmt.Printf("%g\n", ebullicionF - CaF(CongelacionC)) // "180" °F
	//fmt.Printf("%g\n", ebullicionF - CongelacionC) // Error, tipos incompatibles
}