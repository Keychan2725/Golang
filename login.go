package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/gorilla/sessions"
	"database/sql"
    _ "github.com/go-sql-driver/mysql"
    "html/template"
)

var (
    store = sessions.NewCookieStore([]byte("secret-key"))
)

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", homeHandler).Methods("GET")
    r.HandleFunc("/login", loginHandler).Methods("GET")
    r.HandleFunc("/login", loginPostHandler).Methods("POST")
    r.HandleFunc("/logout", logoutHandler).Methods("GET")
    http.Handle("/", r)
    fmt.Println("Server started at localhost:8080")
    http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    session, _ := store.Get(r, "session")
    username, ok := session.Values["username"].(string)
    if !ok {
        username = ""
    }
    t, _ := template.ParseFiles("index.html")
    t.Execute(w, username)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("login.html")
    t.Execute(w, nil)
}

func loginPostHandler(w http.ResponseWriter, r *http.Request) {
    session, _ := store.Get(r, "session")
    r.ParseForm()
    username := r.Form.Get("username")
    password := r.Form.Get("password")
    // Implementasi validasi login (contoh sederhana)
    if username == "user" && password == "pass" {
        session.Values["username"] = username
        session.Save(r, w)
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }
    http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
    session, _ := store.Get(r, "session")
    delete(session.Values, "username")
    session.Save(r, w)
    http.Redirect(w, r, "/", http.StatusSeeOther)
}