package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello, you`ve request: %s\n", r.URL.Path)
	})

	http.ListenAndServe(":80", nil)
}
