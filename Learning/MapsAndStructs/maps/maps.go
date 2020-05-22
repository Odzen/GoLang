package main

import (
	"fmt"
)

//Very valuable and useful data type; when you whant to save one keytype over another value type
// map[keytype]type
func main() {
	/*statePopulations:= make (map[string]int)
	statePopulations = map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}
	*/

	statePopulations := map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}

	fmt.Println(statePopulations)
	fmt.Println(statePopulations["Ohio"])

	statePopulations["Georgia"] = 10310371 // add an element
	fmt.Println(statePopulations["Georgia"])
	fmt.Println(statePopulations)       // the order of the storage we cannot determinate it
	delete(statePopulations, "Georgia") // delete
	fmt.Println(statePopulations)
	fmt.Println(statePopulations["Georgia"]) // although I deleted it, we have the value 0

	// To review if the key was found or not we can use the following

	popTest, ok := statePopulations["Oho"] // pop is the name of the variable associte, and ok is the vab that make us know if that was found, bool
	fmt.Println(popTest, ok)
	_, ok1 := statePopulations["New York"] // if i Just want to know if the key exist
	fmt.Println(ok1)

	fmt.Println(len(statePopulations))

	//If you pass the map as a fuction; have to know that works as slices, is passed by reference
	sp := statePopulations
	delete(sp, "Ohio")
	fmt.Println(sp)
	fmt.Println(len(statePopulations)) // it was delete in both instances
	fmt.Println(len(sp))

	//if exists, i declare the vairable at the same time in a conditona, but  this variable is only define within {} 
	if popu, ok := statePopulations["Florida"]; ok {
		fmt.Println(popu)
	}

}
