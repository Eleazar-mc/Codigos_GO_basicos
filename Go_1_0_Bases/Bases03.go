package main

import "fmt"

func main() {
	var a = 5
	b := 555

	fmt.Println(a, &a) // &a devuelve la dirección de la variable
	fmt.Println(b, &b)
	c := &b
	fmt.Println(c, *c) // *c provee un apuntador a la variable

	var ip *int // ip es un puntero a un entero
	fmt.Println(ip, &ip)

	ip = &a // ip apunta a a
	fmt.Println(ip, &ip)
	y := *ip //y toma valor de 5
	fmt.Println(y, &y)

	ptr := new(int)
	fmt.Println(ptr, &ptr, *ptr)
	*ptr = 3
	fmt.Println(ptr, &ptr, *ptr)
}