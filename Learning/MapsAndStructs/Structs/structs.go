package main

import (
	"fmt"
)

//The advantage of the structs is that we can mix different types of data in the same packet

type Doctor struct {
	number     int
	actorName  string
	companions []string
}

func main() {
	aDoctor := Doctor{
		number:    3,
		actorName: "Jon Pertwee",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}

	fmt.Println(aDoctor)
	fmt.Println(aDoctor.actorName, aDoctor.number, aDoctor.companions[1])

	//another form to declarate a struct and demostrate that
	// The structs are not like maps, htey not passes by reference
	//you can reasigned and not change the original, they are independent
	//if we wnat that they'd point to the same, you can use pointers and pass by reference

	aPatient := struct{ name string }{name: "Juan Sebastian"} // struct{key}{inicializer}
	fmt.Println(aPatient)

	anotherPatience := aPatient
	anotherPatience.name = "Valentina"
	fmt.Println(aPatient)
	fmt.Println(anotherPatience)

	// WE DONT HAVE HERITANCE IN GO, IT DOESNT FOLLOW THAT PRINCIPALE OF THE POO

}
