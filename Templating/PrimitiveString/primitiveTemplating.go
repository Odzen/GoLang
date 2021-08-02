package main

import "fmt"

func main(){
	name:="Juan Velasquez"

	//String in a html doc
	//Merging the name with the html doc
	tpl:= `
		<!DOCTYPE html>
		<html lang="eng">
		<head>
		<meta charset="UTF-8">
		<title>Hello World!</title>
		</head>
		<body>
		<h1>`+name+` </h1>
	    </body>
		</html>
		`

	fmt.Println(tpl)
}
