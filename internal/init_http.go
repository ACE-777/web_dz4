package internal

import (
	"log"
	"net/http"
)

func StartServer() (err error) {
	http.HandleFunc("/login/", login)
	http.HandleFunc("/proceed", proceed)

	log.Println("server run on 9000 port")
	err = http.ListenAndServe("localhost:9000", nil)

	return
}
