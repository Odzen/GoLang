package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	//Step 1: Passed in the data
	//Step 2: Took the value of the data and assing it to a variable, $wisdom in the file .gohtml
	//Step 3: I'm printing that variable out.

	// This time we passed the value of type String
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", `Release self-focus; embrace other-focus.`)
	if err != nil {
		log.Fatalln(err)
	}
}