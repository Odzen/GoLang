package main

import(
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl= template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	//Data
	xs:= []string{"zero", "one", "two","three", "four","five"}

	//Execute Template
	err:= tpl.Execute(os.Stdout, xs)

	if err != nil{
		log.Fatalln(err)
	}
}
