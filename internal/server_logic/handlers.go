package server_logic

import (
	"fmt"
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
	Checkboxone    string
	Checkboxtwo    string
	Volume         string
	Start          string
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method must be GET", http.StatusNotFound)
	}
	if err := tmpl.ExecuteTemplate(w, "login.html", nil); err != nil {
		log.Fatalf("can not execute template login: %v", err)
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
	checkboxone := r.FormValue("checkboxone")
	checkboxtwo := r.FormValue("checkboxtwo")
	start := r.FormValue("start")
	volume := r.FormValue("volume")
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
		Checkboxone:    checkboxone,
		Checkboxtwo:    checkboxtwo,
		Volume:         volume,
		Start:          start,
	}

	if err = tmpl.ExecuteTemplate(w, "processor.html", cred); err != nil {
		log.Fatalf("can not execute template proceed: %v", err)
	}

	w.WriteHeader(http.StatusOK)

	return
}
