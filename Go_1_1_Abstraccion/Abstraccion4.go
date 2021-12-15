package main

import "fmt"

type Tank struct {
	level int
}

func main() {
	t1 := Tank{}
	t2 := Tank{}
	t1.level = 9
	t2.level = 47
	fmt.Println("1: t1.level: ", t1.level, ", t2.level: ", t2.level)
	t1 = t2 // No es un aliasing
	fmt.Println("2: t1.level: ", t1.level, ", t2.level: ", t2.level)
	t1.level = 27 // Independiente de t2.level
	fmt.Println("3: t1.level: ", t1.level, ", t2.level: ", t2.level)
}