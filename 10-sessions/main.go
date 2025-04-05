package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	key   = []byte("session-key")
	store = sessions.NewCookieStore(key) // 存储session的结构
)

func login(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie1")

	session.Values["authed"] = true
	session.Save(r, w)
}

func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie1")

	session.Values["authed"] = false
	session.Save(r, w)
}

func secret(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie1")

	if auth, ok := session.Values["authed"].(bool); ok != true || !auth {
		fmt.Fprintf(w, "forbidden", http.StatusForbidden)
		return
	}

	fmt.Fprintf(w, "haha, you found it!")
}

func main() {
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/secret", secret)

	http.ListenAndServe(":80", nil)
}
