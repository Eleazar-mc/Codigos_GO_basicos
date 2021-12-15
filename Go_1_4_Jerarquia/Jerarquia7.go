package main

import "fmt"

type IGrain interface {
	toString() string
}

type Grain struct {}
func(g Grain) toString() string  {
	return "Grain"
}

type Wheat struct {}
func(w Wheat) toString() string  {
	return "Wheat"
}

type IMill interface {
	process() IGrain
	process2() interface{}
}

type Mill struct {}
func(m Mill) process() IGrain  {
	return Grain{}
}
func(m Mill) process2() float64  {
	return 7.7
}

type WheatMill struct {}
func(w WheatMill) process() IGrain  {
	return Wheat{}
}
func(w WheatMill) process2() string  {
	return "x"
}

func main()  {
	m := Mill{}
	g := m.process()
	fmt.Println(g.toString())

	m = Mill(WheatMill{}) // Equivalencia a upcasting, pérdida del acceso al método "especializado"
	g = m.process()
	fmt.Println(g.toString())
}