package main

import (
	"fmt"
	"strconv"
)

func main() {
	//A simple converson doing casting // 1-2 lines
	var i int = 42
	fmt.Printf("%v, %T\n", i, i)

	var j float32
	j = float32(i) //cast
	fmt.Printf("%v, %T\n ", j, j)

	//A conversion to String but with UNicode // 3-4 lines
	var x int = 92
	fmt.Printf("%v, %T\n", x, x)

	var y string
	y = string(x) //It's looking for the UNICODE of 34
	fmt.Printf("%v, %T\n ", y, y)

	// A rigth conversion to String 5-6 lines / using the library strconv

	var d int = 31
	fmt.Printf("%v, %T\n", d, d)

	var t string
	t = strconv.Itoa(d) //It's looking for the UNICODE of 34
	fmt.Printf("%v, %T\n ", t, t)

}
