package main

import (
	"fmt"
	"github.com/WalterPaes/BuscaCEP/pkg/adapter"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/cep/", handler)
	log.Println("Server is Ok!")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func handler(writer http.ResponseWriter, request *http.Request) {

	var status int
	var result []byte

	// Checking if request method is GET
	if request.Method == "GET" {
		// Get the PostalCode number of url path
		c := strings.TrimPrefix(request.URL.Path, "/cep/")

		// Init the search
		result, status = adapter.CepSearch(c)
	} else {
		result = []byte("Invalid Request Method!")
		status = http.StatusMethodNotAllowed
	}

	writer.WriteHeader(status)
	writer.Header().Set("Content-Type", "application/json")
	fmt.Fprint(writer, string(result))
}
