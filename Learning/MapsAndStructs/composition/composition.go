package main

import (
	"fmt"
)

// WE DONT HAVE HERITANCE IN GO, IT DOESNT FOLLOW THAT PRINCIPALE OF THE POO , instead
// go haves something called "composition" and embedding

type Animal struct {
	Name   string
	Origin string
}

//In normal pro lang the bird would be an animal, and heridate the camps
type Bird struct {
	Animal   // embebedd
	SpeedKPH float32
	CanFly   bool
}

// In the composition we can say better that the Bird has an Animal, it has the caratetisthics and the camps
// But it does not heridate, they are differents structures
// In order to use data in heridate form, we use interfaces
func main() {
	/*
		b := Bird{}
		b.Name = "Emu"
		b.Origin = "Australia"
		b.SpeedKPH = 48
		b.CanFly = false
	*/
	b := Bird{
		Animal:   Animal{Name: "Emu", Origin: "Australia"},
		SpeedKPH: 30,
		CanFly:   false,
	}
	fmt.Println(b)
	fmt.Println(b.Name)

	//Generally we you want to describe common behavior its better to use interfaces

}
