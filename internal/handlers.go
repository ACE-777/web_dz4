package internal

import (
	"html/template"
	"log"
	"net/http"
)

type Cred struct {
	Username string
	Password string
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method must be GET", http.StatusNotFound)
	}

	file, err := template.ParseFiles("cmd/templates/login.html")
	if err != nil {
		log.Printf("Can not render login page: %v", err)
	}

	err = file.Execute(w, nil)
	if err != nil {
		log.Printf("Can not execute template: %v", err)
	}

	w.WriteHeader(http.StatusOK)

	return
}

func proceed(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method must be POST", http.StatusNotFound)
	}
	username := r.FormValue("username")
	password := r.FormValue("password")
	//decode := json.NewDecoder(r.Body)
	//decode.DisallowUnknownFields()
	//
	//var cred Cred
	//
	//if err := decode.Decode(&cred); err != nil {
	//	http.Error(w, "can not decode input data", http.StatusBadRequest)
	//}
	var cred = Cred{Username: username, Password: password}

	file, err := template.ParseFiles("cmd/templates/processor.html")
	if err != nil {
		log.Printf("Can not render login page: %v", err)
	}

	err = file.Execute(w, cred)
	if err != nil {
		log.Printf("Can not execute template: %v", err)
	}
	http.RedirectHandler("/auth", http.StatusOK)
	w.WriteHeader(http.StatusOK)

	return
}
