package main

import "fmt"

type Letter struct {
	c byte // byte es un alias para uint8 (8 bits = 0 a 255)
	d rune // rune es un alias para int32 (32 bits = -2^31 to 2^31 -1)
}

func f(y *Letter) { // En Go no existen los métodos estáticos
	y.c = 'z'
}

func main() {
	x := Letter{}
	x.c = 'a'
	fmt.Println("1: x.c: ", x.c)
	f(&x)
	fmt.Println("2: x.c: ", x.c)
}