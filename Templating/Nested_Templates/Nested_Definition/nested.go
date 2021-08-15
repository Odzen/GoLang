package main

import (
	"text/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl= template.Must(template.ParseGlob("Templates/*.gohtml"))
}

func main() {

	//Execute the main template, in this case, the index.gohtml
	err:= tpl.ExecuteTemplate(os.Stdout, "index.gohtml",42)
	if err != nil {
		log.Fatalln(err)
	}
}
