package main

import (
	"fmt"
)

func main() {
	/*
		a := [3]int{1, 2, 3}
		b := &a[0] // b is a pointer that points to the storage of the fisrt value of the array a
		c := &a[1]
		fmt.Printf("%v %p %p\n", a, b, c) it should print [1 2 3] 0xc00011c140 0xc00011c148
	*/

	var ms *myStruct
	fmt.Println(ms) // always initalizes with the value nil
	//ms = &myStruct{foo: 42} // we create a var to point to the struct, if we want to initialize
	ms = new(myStruct) //it shuould print &{0}, because with new we cannot initialize the value
	/*
		(*ms).foo = 42     // the () its because it have to dereferece first the ms
		fmt.Println(ms)
		fmt.Println((*ms).foo)
	*/
	//We can do it without the (), the compiler can understand
	ms.foo = 42 // the () its because it have to dereferece first the ms
	fmt.Println(ms)
	fmt.Println(ms.foo)
}

type myStruct struct {
	foo int
}
