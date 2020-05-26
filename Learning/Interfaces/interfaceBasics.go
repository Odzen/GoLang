package main

import (
	"fmt"
)

func main() {

}

//Innterface, they dont describe data, they describe behaviours
type Writter interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}
