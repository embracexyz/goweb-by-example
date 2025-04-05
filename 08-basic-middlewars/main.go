package main

import (
	"fmt"
	"net/http"
)

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		f(w, r)
	}
}

func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "bar heihei")
}

func main() {
	http.HandleFunc("/bar", logging(bar))
	http.ListenAndServe(":80", nil)
}
