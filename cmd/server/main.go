package main

import (
	"log"
	"web_dz4/internal"
)

func main() {
	err := internal.StartServer()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
