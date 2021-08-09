package main

import (
	"log"
	"os"
	"text/template"
)

// Package level scope, the variable tpl it's accessible throughout this package
var tpl *template.Template

// func init() = Initializes the program, it runs once when the program is starting up
// Must(*template, err) and ParseGlob returns a *template and an error. So it fits
// The function Must handle the error for us, so we use to not to handle it.
func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	//In the main I execute my templates...
	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "vespa.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
