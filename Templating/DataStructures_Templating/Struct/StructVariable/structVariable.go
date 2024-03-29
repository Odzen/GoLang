package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct{
	Name string
	Motto string
}

func init()  {
	tpl= template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	//Instantiate the structure
	buddha := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	err:= tpl.Execute(os.Stdout,buddha)

	if err != nil {
		log.Fatalln(err)
	}


}
