package main

import "fmt"

type Use interface {
	f()
	g()
}

type Useful struct {}
func(u Useful) f()  {
	fmt.Println("Useful.f()")
}
func(u Useful) g()  {
	fmt.Println("Useful.g()")
}

type MoreUseful struct {}
func(m MoreUseful) f()  {
	fmt.Println("MoreUseful.f()")
}
func(m MoreUseful) g()  {
	fmt.Println("MoreUseful.g()")
}
func(m MoreUseful) u()  {
	fmt.Println("MoreUseful.u()")
}
func(m MoreUseful) v()  {
	fmt.Println("MoreUseful.v()")
}
func(m MoreUseful) w()  {
	fmt.Println("MoreUseful.w()")
}

func main()  {
	x := [2]Use{Useful{},MoreUseful{}}
	x[0].f()
	x[1].g()

	y := [2]MoreUseful{MoreUseful(Useful{}),MoreUseful{}} // No hace uso de la interfaz
	y[0].f()
	y[1].g()
	y[1].u()
	Useful(y[0]).f()

}