package main

import (
	"fmt"
)

func main() {
	s := sum(1, 2, 3, 4, 5)
	fmt.Println("The sume is", s)
	//sum("The sum is", 1, 2, 3, 4, 5)
}

//func to return sintax // func name(parameters...) type_of_return
func sum(values ...int) int { // the notacion ... means that it have to take all the parameters and make and slice with them
	fmt.Println(values) // now is a slide
	result := 0
	for _, v := range values {
		result += v
	}
	return result
}

/*
//We can say the same but with some sintatic sugar
func sum(values ...int) (result int) { // Im saying before the defintion that the value that it must to return has to be the var result
	fmt.Println(values) // now is a slide
	result := 0
	for _, v := range values {
		result += v
	}
	return // we dont need to say which value it has to return because we already put it above
}
*/
/*
//We can also return pointers
func sum(values ...int) *int { // it returns a pointer -> type_int
	fmt.Println(values) // now is a slide
	result := 0
	for _, v := range values {
		result += v
	}
	return &result
}
*/

/*
//If you use this kind of parameters it can only be one, and it has to be at the end
func sum(msg string, values ...int) { // the notacion ... means that it have to take all the parameters and make and slice with them
	fmt.Println(values) // now is a slide
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println(msg, result)
}
*/
