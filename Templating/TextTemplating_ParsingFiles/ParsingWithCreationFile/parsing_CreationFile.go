package main

import (
	"log"
	"os"
	"text/template"
)

func main() {

	//Step 1, parse files
	tpl, err := template.ParseFiles("fileHtml.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	//Step 2, Create the file in which we will put the parse
	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatalln("error creating file", err)
	}
	defer nf.Close()

	//Step 3. Execute, and put in the file. Since nf is a file.
	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatalln(err)
	}
}