package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//Run in LIFO // last in - first out order

func main() {
	/*defer
	fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)

	a:= "start"
	defer fmt.Println(a)
	a="end"

	//it must print start because, allthought it prints after the main body, it takes the first or the closest definition
	*/

	res, err := http.Get("https://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close() // its allows the open of the resource and the close to the resource
	//right next to each other, because the defer allows to make the action after the exectution of the main
	//it is not the best for a loop
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)

}
