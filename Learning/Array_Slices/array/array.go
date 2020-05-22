package main

import (
	"fmt"
)

func main() {
	grades := [3]int{97, 85, 93}
	// grades := [...]int {97,52}
	var students [3]string
	fmt.Printf("Grades: %v\n", grades)
	fmt.Printf("Students: %v\n", students)
	students[0] = "Lisa"
	students[1] = "Juan"
	students[2] = "David"
	fmt.Printf("Students: %v ,Lenght: %v\n", students, len(students))

	//Matrices

	var identityMatrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
	fmt.Println(identityMatrix)

	//Function like variables

	//The value b is assigned to a copy of the array a
	a := [...]int{1, 2, 3}
	b := a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)

	//The t point on the same data that i has

	t := [...]int{1, 2, 3}
	i := &t
	i[1] = 5
	fmt.Println(t)
	fmt.Println(i)

}
