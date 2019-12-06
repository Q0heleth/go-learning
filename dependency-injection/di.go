package main

import (
	"fmt"
	"io"
	"math"
	_ "math"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}
func MygreetHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}
func main() {
	fmt.Println(math.Pi)
	http.ListenAndServe(":5000", http.HandlerFunc(MygreetHandler))
}
