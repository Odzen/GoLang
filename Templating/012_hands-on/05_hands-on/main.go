package main

import (
	"log"
	"os"
	"text/template"
)

type restaurants struct {
	TimeDay string; // Morning, Arternoon, Night

}

type


var tpl *template.Template

func init()  {
	tpl= template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	err:= tpl.Execute(os.Stdout, 0)

	if err != nil{
		log.Fatalln(err)
	}
}
