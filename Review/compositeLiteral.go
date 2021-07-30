package main

import "fmt"

//Go is everything about types, you can create your own types
type hotdog int

//Struct:
type person struct {
	//Fields, starts with UpperCase if I want them to be visible outside the package and lowcase only inside
	fname string // Inside
	Lname string // Outside
}

//Functions
//func (receiver) identifier(parameters) (returns) { <code> }
func (p person) speak() {
	fmt.Println(p.fname, p.Lname, `says, "Good morning, James."`)
}

//For SecretAgent embedded types
func (sa secretAgent) speak() {
	fmt.Println(sa.fname, sa.Lname, `says, "Good morning, Agent."`)
}

//Compostions embedded types
type secretAgent struct {
	person
	licenseToKill bool
}

//Interfaces, defines functionallity and allow polymorphism

type human interface {
	speak()
}

func main() {
	//Composite Literal --> type{...}
	var t hotdog
	t = 7
	fmt.Println(t)

	xi := []int{2, 5, 3, 4}
	fmt.Println(xi)

	//Map
	m := map[string]int{
		"Todd": 45,
		"Job":  42,
	}
	fmt.Println(m)

	//Structs using
	p1 := person{
		"Mr",
		"Juanse",
	}

	fmt.Println(p1)

	//Functions
	p1.speak()

	//Composition embedded types
	secretAgent1 := secretAgent{
		person{
			"James",
			"Bond",
		},
		true,
	}

	//To access to the method for a SecretAgent
	secretAgent1.speak()
	//To acces to the method for Person, calling inner type
	secretAgent1.person.speak()

}
