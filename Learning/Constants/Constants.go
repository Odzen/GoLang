package main

import (
	"fmt"
)

//it cant be shawdoed 
const (
	_ = iota // ignore the fist value 
	x = iota
	y = iota
	z = iota
)

func main() {
	const myConst int = 42
	//myConst=5  ERROR it cant be done
	//const myConst float64=math.Sin(1.57)ERROR it has to be assginable at compile time, because requeres the funtion to execute
	fmt.Printf("%v, %T\n", myConst, myConst)

	const (
		a int     = 14
		b string  = "foo"
		c float32 = 3.14
		d bool    = true
	)
	fmt.Printf("%v, %v, %v,%v\n", a, b, c, d)
	fmt.Printf("%v, %v, %v", x, y, z)
}
