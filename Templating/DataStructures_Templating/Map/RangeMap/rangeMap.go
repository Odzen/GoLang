package main

import (
	"os"
	"log"
	"text/template"
)

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}
func main() {
	//In this case it will print the values but not the key, because it is not evaluated in the .gohtml

	sages:= map[string]string{
		"India":    "Gandhi",
		"America":  "MLK",
		"Meditate": "Buddha",
		"Love":     "Jesus",
		"Prophet":  "Muhammad",
	}

	err:= tpl.Execute(os.Stdout, sages)
	if err!= nil{
		log.Fatalln(err)
	}
}
