package main

import "fmt"

type DataOnly struct {
	i int
	d float64
	b bool
}

func main() {
	d := DataOnly{}
	fmt.Println(d.b, d.d, d.i)
	d.i = 100
	fmt.Println(d.b, d.d, d.i)
}