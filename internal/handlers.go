package internal

import (
	"fmt"
	"html/template"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"
)

type Cred struct {
	Username       string
	Password       string
	TextareaInput  string
	OneRadio       string
	SelectInput    string
	FileInput      string
	NotHiddenInput string
	HiddenInput    string
	FileInputSize  string
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
	var (
		fileName  string
		fileSize  string
		handler   *multipart.FileHeader
		fileInput multipart.File
		err       error
	)

	if r.Method != http.MethodPost {
		http.Error(w, "method must be POST", http.StatusNotFound)
	}
	username := r.FormValue("username")
	password := r.FormValue("password")
	textareaInput := r.FormValue("textareainput")
	oneRadio := r.FormValue("oneradio")
	selectinput := r.FormValue("selectinput")
	nothiddeninput := r.FormValue("titlenothidden")
	hiddeninput := r.FormValue("postId")
	fileInput, handler, err = r.FormFile("file")
	if err != nil {
		if err == http.ErrMissingFile {
			fileName = ""
			fileSize = "0"
		} else {
			http.Error(w, fmt.Sprintf("Ошибка при получении файла: %s", err), http.StatusBadRequest)
			return
		}
	} else {
		fileName = handler.Filename
		fileSize = strconv.Itoa(int(handler.Size))
		defer fileInput.Close()
	}

	var cred = Cred{
		Username:       username,
		Password:       password,
		TextareaInput:  textareaInput,
		OneRadio:       oneRadio,
		SelectInput:    selectinput,
		FileInput:      fileName,
		FileInputSize:  fileSize,
		NotHiddenInput: nothiddeninput,
		HiddenInput:    hiddeninput,
	}

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
