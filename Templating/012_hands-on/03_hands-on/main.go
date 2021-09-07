package main

import (
	"log"
	"os"
	"text/template"
)


type calHotel struct{
	Name string
	Address string
	City string
	Zip int
}

type region struct {
	Region string // Can be Southern, Central, Northern
	Hotels []calHotel // array to hold a unlimited number of hotels
}

var tpl *template.Template

func init()  {
	tpl= template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	regiones:=region{
		Region: "Southern",
		Hotels: []calHotel{
			{
				"Name", "Address", "City", 1234,
			},

			{
				"Name2", "Address", "City", 1235,
			},
		},
	}
	err := tpl.Execute(os.Stdout, regiones)
	if err != nil {
		log.Fatalln(err)
	}
}
