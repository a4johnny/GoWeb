package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	// go get github.com/gorilla/mux
	// go install github.com/gorilla/mux
)

// ex :  http://localhost:3000/?key=abc // output > Url Param 'key' is: abc //參數的數量與位置不會影響
// func main() {
// 	http.HandleFunc("/", handler)
// 	http.ListenAndServe(":3000", nil)
// }

// func handler(w http.ResponseWriter, r *http.Request) {
// 	keys, ok := r.URL.Query()["key"]
// 	if !ok || len(keys[0]) < 1 {
// 		fmt.Fprint(w, "Url Param 'key' is missing")
// 		return
// 	}

// 	key := keys[0]
// 	fmt.Fprint(w, "Url Param 'key' is : "+string(key))
// }

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You have requested the book: %s on page %s\n", title, page)
	})

	http.ListenAndServe("140.123.97.81:80", r)
}
