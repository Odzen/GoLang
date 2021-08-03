package main

import (
	"log"
	"os"
	"text/template"
)

func main() {

	//STEP 1. PARSE FILES

	//ParseFiles takes 0 or more strings, separate with ,
	//ParseFiles gives and * Template and an error
	tpl,err:= template.ParseFiles("fileHtml.gohtml")

	if err!= nil{
		log.Fatalln(err)
	}

	//STEP 2. EXECUTE FILES
	//The Method Execute takes a Writer and data
	err=tpl.Execute(os.Stdout, nil)

	if err!= nil{
		log.Fatalln(err)
	}

}
