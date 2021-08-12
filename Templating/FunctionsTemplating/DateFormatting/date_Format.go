package main

import (
	"time"
	"text/template"
	"os"
	"log"
)

var tpl *template.Template

func init(){
	tpl= template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

// The function that formats the time
// The Format method takes a layout string and gives a string
// Just different time standars to show to the program how can I print my time
func monthDayYear(t time.Time) string {
	//Because of the format, when I want the month I put 1, where I want the day I put the 2, where I want the year I put 6
	return t.Format("02-01-2006")
}

// The FuncMap, the type map that we will give to our template
var fm=template.FuncMap{
	"fdateMDY":monthDayYear,
}

func main() {
	// Executing passing the current time as data {{.}}
	err:= tpl.ExecuteTemplate(os.Stdout,"tpl.gohtml", time.Now())

	if err != nil{
		log.Fatalln(err)
	}
}
