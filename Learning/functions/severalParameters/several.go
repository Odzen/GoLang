package main

import (
	"fmt"
)

func main() {
	d, err := divide(5.0, 0.0)
	if err != nil {
		fmt.Println(err)
		return // return in the main function its going to exit the app
	}
	fmt.Println(d)

}

func divide(a, b float64) (float64, error) { // we return the error in order to manage it, and no stop the program
	//instead of using panic that stops the program
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero") //  we return the errorObject
		//panic("Cannot provide zero as second value")
	}
	return a / b, nil
}
