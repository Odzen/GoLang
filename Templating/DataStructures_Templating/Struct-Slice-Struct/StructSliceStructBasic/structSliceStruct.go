package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

type items struct {
	Wisdom    []sage
	Transport []car
}

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	//Structs

	//Struct sages
	sageBuddha := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	sageGandhi := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}

	sageMKing := sage{
		Name:  "Martin Luther King",
		Motto: "Hatred never ceases with hatred but with love alone is healed.",
	}

	//Struct Cars
	carFord := car{
		Manufacturer: "Ford",
		Model:        "F150",
		Doors:        2,
	}

	carToyota := car{
		Manufacturer: "Toyota",
		Model:        "Corolla",
		Doors:        4,
	}

	// Slides of structs
	sages := []sage{sageBuddha, sageGandhi, sageMKing}
	cars := []car{carFord, carToyota}

	// data is a struct of slides (Wisdom and Transport), in at the same time, each item of those slides are structures
	data := items{
		Wisdom:    sages,
		Transport: cars,
	}

	//Execute
	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}