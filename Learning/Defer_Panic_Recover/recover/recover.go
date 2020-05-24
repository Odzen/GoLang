package main

import (
	"fmt"
	"log"
)

/*
	//Anonymous function, doesnt have a name so nothing can call this, it is define and we can call it exactly 1 time
	//recover=its gonna return nil if the app isnt panic
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
		}
	}() // I call it here

	panic("something bad")
	fmt.Println("end")
*/

func main() {
	fmt.Println("start")
	panicker()
	fmt.Println("end")
}

func panicker() {
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err) // we print the error but in the main function the execution continues
			//panic(err) repanic
			// thats why it print the "end"
			// if i recover the error it means that youre gonna deal with it
			// if you dont want to deal with it, and want to stop the execution of the main
			// you can put repanic
		}
	}()
	panic("something bad")
	fmt.Println("done")
}
