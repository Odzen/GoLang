package main

import (
	"fmt"
)

/*
Pasing a pointers is much more efficient because if you pass by value you force the program to make a copy
every single time that it calls the function
of course the difference it can be noticed in large data structures
Remember that the slices and structures always passing their values by reference
the arrays and the maps always by value
*/

//PASSING BY REFERENCE WITH POINTERS, the value name change in both functions the general and the one who is
// in the function, since is passed by reference it is the same value and not a copy
func main() {
	greeting := "Hello"
	name := "Stacey"
	sayGreeting(&greeting, &name) // Im passing the values by reference
	fmt.Println(name)             // the name is ted since it was change in the funtions called before
}

func sayGreeting(greeting, name *string) { // *string to say that greeting and name are pointers -> string
	fmt.Println(*greeting, *name) // dereference to pull out the values
	*name = "Ted"                 // dereference to pull out the value and changing, here Im changing also the general value
	fmt.Println(*name)            // the name is ted
}

/*
//PASSING BY VALUE, the value name change in the functions but since is passed by valor
// it makes a copy and only changes in the func that is called but not in the main value
func main() {
	greeting := "Hello"
	name := "Stacey"
	sayGreeting(greeting, name)
	fmt.Println(name) // the name is still stacey
}

func sayGreeting(greeting, name string) {
	fmt.Println(greeting, name)
	name = "Ted"
	fmt.Println(name) // the name is ted
}
*/
