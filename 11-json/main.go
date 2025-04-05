package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type user struct {
	Name string `json:Name`
	Age  int    `json:Age`
}

func main() {
	http.HandleFunc("/encode", func(w http.ResponseWriter, r *http.Request) {
		u1 := user{
			Name: "lihua",
			Age:  18,
		}
		json.NewEncoder(w).Encode(u1)
	})

	http.HandleFunc("/decode", func(w http.ResponseWriter, r *http.Request) {
		var u1 user
		json.NewDecoder(r.Body).Decode(&u1)
		fmt.Printf("#v\n", u1)
	})

	http.ListenAndServe(":80", nil)
}
