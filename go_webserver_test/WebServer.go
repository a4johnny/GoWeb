package main

import (
	"fmt"
	"log"
	"net/http"
)

var count int

func main() {
	http.HandleFunc("/", handler) //like route 每個request呼叫handler這個函式
	http.HandleFunc("/count/", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil)) // 監聽local port 8000 並記錄log
}

func handler(w http.ResponseWriter, r *http.Request) { //(類似 controller)
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	count += 1
	fmt.Fprintf(w, "Count %d\n", count)
}
