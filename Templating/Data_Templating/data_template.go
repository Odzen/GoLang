package main

import(
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init(){
	 tpl= template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	//So in this case, we pass the data, it is not nil, it is 42
	// It will print the file, with the data 42 which is an int
	err:= tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml",42)
	if err!= nil{
		log.Fatalln(err)
	}
}
