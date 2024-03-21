package main

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	//http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler))
	err := http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler))
	if err != nil {
		fmt.Println(err)
	}
}
