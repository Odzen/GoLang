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

}
