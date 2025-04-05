package main

import (
	"fmt"
	"net/http"
	"time"
)

type middleware func(http.HandlerFunc) http.HandlerFunc

func logging() middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			t := time.Now()
			defer func() {
				fmt.Println(r.URL.Path, time.Since(t))
			}()

			f(w, r)
		}
	}
}

func Method(method string) middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if r.Method != method {
				fmt.Fprintf(w, "unsupprot method, only GET for now!")
				return
			}

			f(w, r)
		}
	}
}

func Chain(f http.HandlerFunc, middlewares ...middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello CS")
}

func main() {
	http.HandleFunc("/", Chain(hello, logging(), Method("GET")))
	http.ListenAndServe(":80", nil)
}
