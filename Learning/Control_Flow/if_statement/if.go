package main

import (
	"fmt"
)

func main() {
	//thats a sintax error
	/*
		if true
			fmt.Println("")
	*/
	if true {
		fmt.Println("pass")
	}

	number := 50
	guess := -2

	if guess < 1 || guess > 100 {
		fmt.Println("The number must be 1-100")
	} else {
		if guess < number {
			fmt.Println("To low")
		}
		if guess > number {
			fmt.Println("To high")
		}
		if guess == number {
			fmt.Println("You got it")
		}
		fmt.Println(number <= guess, number >= guess, number != guess)
	}
}
