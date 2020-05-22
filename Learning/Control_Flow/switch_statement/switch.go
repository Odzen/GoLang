package main

import (
	"fmt"
)

func main() {
	switch 5 {
	case 1, 5, 10:
		fmt.Println("one")
	case 2, 3, 6:
		fmt.Println("two")

	default:
		fmt.Println("not one or two")

	}

	//In go it is not neccesary use the breaks; they are implicit
	i := 10
	switch {
	case i <= 10:
		fmt.Println("less than or equal to ten")
		fallthrough /// to continue executing, ypu dont use it very often
	case i <= 20:
		fmt.Println("less than or equal to twenty")
	default:
		fmt.Println("greater than 20")

	}

	// Type swicth

	var x interface{} = 5 // is type interface so it could be any thing
	//var y interface{} = [3]int{}
	switch x.(type) {
	case int:
		fmt.Println("intx")
		break
		fmt.Println("test")
	case float64:
		fmt.Println("x=float")
	case string:
		fmt.Println("x=string")
	default:
		fmt.Println("x is another type ")

	}

}
