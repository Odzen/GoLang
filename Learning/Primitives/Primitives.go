package main

import (
	"fmt"
)

func main() {
	var n bool = true
	x := 1 == 1
	var y bool
	fmt.Println(n)
	fmt.Println(x, y)

	// Operaciones primitivas
	a := 10 // 1010 binary value for oparation with binaries
	b := 3  // 0011 binary value for oparation with binaries

	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	fmt.Println(a & b)  // 0010
	fmt.Println(a | b)  // 1011
	fmt.Println(a ^ b)  // EXCLUVISE OR // 1001
	fmt.Println(a &^ b) // AND NOT OPERATOR // 0100

	//Floats

	e := 3.14
	g := 13.7e72

	fmt.Println(g, e)

	//var a int = 9
	//var b int8 = 2
	// a+b it isnt posible, you have to do a cast

	//Complex numbers for Data science

	var t complex64 = 1 + 2i // 64 and 128 obligatory // or complex(5,12) real, img
	fmt.Println(t)
	fmt.Printf("%v, %T\n", real(t), real(t)) // take the parts of the number that divides the bits in 2
	fmt.Printf("%v, %T\n", imag(t), imag(t))

	//Strings // works as an array // treats also for bites

	s := "Juanse"
	//fmt.Println("%v, %T\n", s[2], s[2]) it returns a number of bites, so you need to cast the result
	fmt.Println("%v, %T\n", string(s[2])+" Hola ", string(s[2]))

	// Characters
	u := 'c'
	fmt.Println("%v , %T\n", u, u) // it returns an int32 because the characters in goLang are alias to int32

}
