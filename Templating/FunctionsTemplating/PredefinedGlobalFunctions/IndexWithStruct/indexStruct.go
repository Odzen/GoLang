package main

import(
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl= template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	//Data
	xs:= []string{"zero", "one", "two","three", "four","five"}

	//Data, slide into a Struct
	data := struct {
		Words []string
		Lname string
	}{
		xs,
		"McLeod",
	}

	//Execute Template
	err:= tpl.Execute(os.Stdout, data)

	if err != nil{
		log.Fatalln(err)
	}
}
