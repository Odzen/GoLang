package main

import (
	"fmt"
)

func main() {
	/*
		a := 42
		b := a // you copy the data storage in a and assigened in b
		fmt.Println(a, b) it should print 42 42
		a = 27
		fmt.Println(a, b) it should print 27 42
	*/
	var a int = 42
	var b *int = &a    // b->int and it points to a value of a // the b is a pointer
	fmt.Println(a, b)  // it should print the number and the memory location of storage a
	fmt.Println(&a, b) // &= reference operator, it shows the direction
	fmt.Println(a, *b) // *= dereference operator, it shows the value
	// the *type = declares a pointer to a data of that type
	// the *var=pointer= it derefereces; pull the value back out
	a = 27
	fmt.Println(a, *b) // the are both point at the same value
	*b = 14            // we change the value
	fmt.Println(a, *b)
}
