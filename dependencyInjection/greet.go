package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func main() {

	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))

	Greet(os.Stdout, "Raphael")
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}
