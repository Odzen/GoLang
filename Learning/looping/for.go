package main

import (
	"fmt"
)

func main() {
	x := 0 // i can generilice to get the final value
	for x < 2 {
		fmt.Println(x)
		x++
	}
	//another notation
	for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}

	//another

	y := 0
	for {
		fmt.Println(y)
		y++
		if y == 5 {
			break
		}
	}

	//with slices

	s := []int{1, 2, 3}
	for k, v := range s {
		fmt.Println(k, v)
	}

	// with strings

	str := "Hello Go"
	for k, v := range str {
		fmt.Println(k, string(v))
	}
}
