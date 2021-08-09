package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	// So this time we use ParseGlob, this function instead use a pattern, so like a regex to identifiy which files to parse
	// template/* it says all the files into the directory templates
	tpl, err := template.ParseGlob("templates/*")
	if err != nil {
		log.Fatalln(err)
	}

	// Execute the first one
	// Print ONE
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	// Executing the vespa
	// Print VESPA
	err = tpl.ExecuteTemplate(os.Stdout, "vespa.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	// Executing the file number two
	// Print TWO
	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	// Executing the file number one
	// Print ONE
	err = tpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
