package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		sayMessage("Hello go", i)
	}
}

func sayMessage(msg string, idx int) {
	fmt.Println(msg)
	fmt.Println("The value of the message is", idx)
}

/* We can add the type at the end to show that all the arguments are of the same type
func sayMessage(msg1 , msg2, msg3 string) {
	...
}
*/
