package server_logic

import (
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("internal/templates/*.html"))
}

func StartServer() (err error) {
	fs := http.FileServer(http.Dir("internal/static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	http.HandleFunc("/login/", login)
	http.HandleFunc("/proceed", proceed)

	log.Println("server run on 9000 port")
	err = http.ListenAndServe(":9000", nil)

	return
}
