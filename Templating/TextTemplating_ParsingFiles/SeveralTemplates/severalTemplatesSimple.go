package main

import (
	"log"
	"os"
	"text/template"
)

func main() {

	//STEP 1: Parsing the file number one, it returns a template
	tpl, err:= template.ParseFiles("one.gohtml")
	if err != nil{
		log.Fatalln(err)
	}

	//STEP 2: Executing the template and print it
	// Will print ONE, since is the only file parsed
	err=tpl.Execute(os.Stdout, nil)
	if err!=nil{
		log.Fatalln(err)
	}


	// Adding more files to parse into the template tpl
	tpl, err = tpl.ParseFiles("two.gohtml", "vespa.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	// Executing them
	// This time we use the method ExecuteTemplate because this allows me to Execute and specif Template
	// ExecuteTemplate(Where, name of the template, data)
	// This will print VESPA
	err = tpl.ExecuteTemplate(os.Stdout, "vespa.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	// Executing the template two
	// This will Print TWO
	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	// Executing the template one
	// This will print ONE
	err = tpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	// Executing the the first file one and printing it, because it's the first file that we put into the container
	// This will print ONE
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}



}
