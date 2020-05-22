package main

import (
	"fmt"
)

func main() {
	//The slices reffers like a pointer automatically, so if you reassigned them
	// it doesnt make a copy, it reffers to the very same, and you dont have to put pointers

	a := []int{1, 2, 3} // without [...] its a slice
	b := a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))

	// Slices operation it can be done in arrays or slices
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	y := x[:]   // slice of all elements
	c := x[3:]  // slice from the 4th element
	d := x[:6]  // slice first 6 elements
	e := x[3:6] // slice the 4th, 5th, and 6th elements

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	//Make function, the reason why we make the capacity like thata, its because the slices can change over the time
	// you can remove or add elements

	t := make([]int, 3, 100) // creation, elements , capacity
	fmt.Println(t)
	fmt.Printf("LengthT: %v\n", len(t))
	fmt.Printf("CapacityT: %v\n", cap(t))

	// Stacks operations
	r := []int{}
	fmt.Println(r)
	fmt.Printf("LengthR: %v\n", len(r))
	fmt.Printf("CapacityR: %v\n", cap(r))
	r = append(r, 1, 2, 3, 4, 5)
	fmt.Println(r)
	fmt.Printf("LengthR: %v\n", len(r))
	fmt.Printf("CapacityR: %v\n", cap(r))

	//Remove elements
	u := r[1:]
	fmt.Println(u)
	u = r[:len(r)-1]
	fmt.Println(u)

	u = append(r[:2], r[3:]...) // removing tthe middle element and have to use the spread operation
	fmt.Println(u)
	fmt.Println(r) // unexpected behaviur because it maj¿ke refference to the same memory
	// To solve that problem, you have to make a copy of the array with a loop

}
