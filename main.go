package main

import (
	"fmt"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hola Mundo")
}

func main() {
	http.HandleFunc("/webhook", helloWorld)
	http.ListenAndServe(":8080", nil)
}
