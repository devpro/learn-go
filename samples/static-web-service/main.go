package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello World!"))
	})

	error := http.ListenAndServe(":3000", nil)

	if error != nil {
		log.Fatal(error)
	}
}
