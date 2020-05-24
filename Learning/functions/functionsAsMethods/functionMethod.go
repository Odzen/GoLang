package main

import "fmt"

func main() {
	g := greeter{
		greeting: "Hello",
		name:     "Go",
	}
	g.greet()
	fmt.Println("The new name is", g.name)
}

type greeter struct {
	greeting string
	name     string
}

/*
//This is the method, a function that is executed in a non contact
// It means in any type // fun (var type) name () {}
//When its called, it makes a copy of the gretter object and assign the name of g
// Its the same as function, just provide a context that is been executed in a certain type
//THIS IS PASSING BY VALOR
func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
	g.name = "" // IT DOES NOT CHANGE THE GENERAL VALUE, IT CAN BE ACCES TO DE DATA BUT NOT CHANGE IT
}
*/
//THIS IS PASSING BY REFERENCE
func (g *greeter) greet() {
	fmt.Println(g.greeting, g.name)
	g.name = "" // IT CHANGE THE GENERAL VALUE AND CAN ACCESS TO THE DATA
}
