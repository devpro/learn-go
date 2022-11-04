package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		names := request.URL.Query()["name"]
		var name string
		if len(names) == 1 {
			name = names[0]
		}

		accept := request.Header["Accept"]
		if len(accept) == 1 && accept[0] == "application/json" {
			// json output
			values := map[string]string{"name": name}
			encoder := json.NewEncoder(writer)
			err := encoder.Encode(values)
			if err != nil {
				writer.Write([]byte("Error while encoding"))
			}
		} else {
			// raw output
			writer.Write([]byte("Hello " + name))
		}
	})

	error := http.ListenAndServe(":3000", nil)

	if error != nil {
		log.Fatal(error)
	}
}
