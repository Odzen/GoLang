package main

import "fmt"

func main() {
	/*  1 way
	var f func() = func() {
		fmt.Println("F")
	}

	2 way
	f := func() {
		fmt.Println("F")
	}
	f()*/

	//YOU HAVE TO DEFINE BEFORE EXECUTE THE FUNCTIONS IF YO WORK WITH THE FUNCTIONS AS VARIABLES
	var divide func(float64, float64) (float64, error)
	divide = func(a, b float64) (float64, error) {
		if b == 0.0 {
			return 0.0, fmt.Errorf("Cannot be divided by 0")
		} else {
			return a / b, nil
		}
	}
	d, err := divide(5.0, 3.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)
}
