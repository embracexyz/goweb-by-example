package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/book/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "you are read page %s of %s", page, title)
	})

	// server-> router(管理route，route可以在自身通过方法限制schema，host，tls，methods等) -> handleFunc
	// handleFunc + path 定义一组route
	http.ListenAndServe(":80", r) // r 为默认的nil时，表示使用http packge 默认的router
}
