package main

import(
	"log"
	"text/template"
	"os"
	"strings"
)


var tpl *template.Template

type sage struct {
	Name string
	Motto string
}

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

// create a FuncMap to register functions.
// "uc" is what the func will be called in the template
// "uc" is the ToUpper func from package strings
// "ft" is a func I declared
// "ft" slices a string, returning the first three characters
// COMPOSITE LITERAL
var fm = template.FuncMap{
	"uc": strings.ToUpper, // Function from the package string to make all upperCase
	"ft": firstThree, // Function that we create to take the first three letters
}

// The init will change a litte bit, we have to use the function New, and .Funcs to give the FunctionMap
// To the template
func init() {
	//You have to give to the template the Functions before you Parsed, Always
	//So to get *template, to call Funcs, so the only one that gives that without Parsing
	// Is the function New(name String) *template, that creates a *template and gives it a name
	// I give no name, because I don't want to parse it, I just want to give it the functions with .Funcs
	// And then Parse it
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("index.gohtml"))
}

// It gets the first three letters of any string
func firstThree(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 3 {
		s = s[:3]
	}
	return s
}

func main() {

	b := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	g := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}

	m := sage{
		Name:  "Martin Luther King",
		Motto: "Hatred never ceases with hatred but with love alone is healed.",
	}

	f := car{
		Manufacturer: "Ford",
		Model:        "F150",
		Doors:        2,
	}

	c := car{
		Manufacturer: "Toyota",
		Model:        "Corolla",
		Doors:        4,
	}

	sages := []sage{b, g, m}
	cars := []car{f, c}

	data := struct {
		Wisdom    []sage
		Transport []car
	}{
		sages,
		cars,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
