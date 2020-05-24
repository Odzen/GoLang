package main

import "net/http"

// WE DONT HAVE EXCEPTIONS IN GO, so we have panic
// the panic only use in unrecoverable events beacuse it stops the app
func main() {
	/*
		a, b := 1, 0
		ans := a / b
		fmt.Println(ans)

		It should print the following
		panic: runtime error: integer divide by zero
	*/

	/* it doesnt reach the line end
	fmt.Println("start")
	panic("bad")
	fmt.Println("end")
	*/
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello go"))
	})
	err := http.ListenAndServe(":8080,nil", nil)
	if err != nil {
		panic(err.Error())
	}

}
