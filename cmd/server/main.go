package main

import (
	"fmt"
	"net/http"
)

func handlers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	http.HandleFunc("/", handlers)
	http.ListenAndServe(":1488", nil)
}
