package main

import (
	"fmt"
	"net/http"
)

func headlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1> fist </h1>")
}

func main() {
	http.HandleFunc("/", headlerFunc) //要求 request
	http.ListenAndServe(":3000", nil)
}
