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
				log.Print("Impossible to write json output in response ", err)
			}
		} else {
			_, err := writer.Write([]byte("Hello " + name))
			if err != nil {
				writer.WriteHeader(500)
				_, err := writer.Write([]byte("Error while writing response"))
				if err != nil {
					log.Print("Impossible to write error in response")
				}
			}
		}
	})

	error := http.ListenAndServe(":3000", nil)

	if error != nil {
		log.Fatal(error)
	}
}
