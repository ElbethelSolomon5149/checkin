package main

import (
	"html/template"
	"net/http"
	// import menu hadler
)

// connect database - previleges?

// handlefunc, start server
var tmpl = template.Must(template.ParseGlob("/UI/templates/*"))

func index(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.layout", nil)
}
func book(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "book.layout", nil)
}
func checkin(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "checkin.html", nil)
}
func main() {
	fs := http.FileServer(http.Dir("/UI/assets"))
	mux := http.NewServeMux()
	mux.Handle("/UI/assets/", http.StripPrefix("/UI/assets/", fs))
	mux.HandleFunc("/", index)
	mux.HandleFunc("/book", book)
	mux.HandleFunc("/checkin", book)
	http.ListenAndServe(":8080", mux)
}
