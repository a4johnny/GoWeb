package main //day 11

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/num/", add)
	if err := http.ListenAndServe(":3000", nil); err != nil {
		panic(err)
	}
}

func add(w http.ResponseWriter, r *http.Request) {
	a := strings.Split(r.URL.Path, "/")
	_, _, s := a[0], a[1], a[2]
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	} else {
		fmt.Fprintf(w, "n + 100 = %d\n", i+100)
	}
}
