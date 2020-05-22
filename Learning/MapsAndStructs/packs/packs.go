package main

import (
	"fmt"
	"reflect"
)

//I can do tags, for example to put restriccions in the camps, you have to import reflect to make that
type Animal struct {
	Name   string `required max:"100"`
	Origin string
}

func main() {
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}
