package main

import (
	"text/template"
	"log"
	"os"
)

var tpl *template.Template

func init()  {
	tpl:= template.Must(template.ParseFiles("table.csv"))
}

func main() {

}
