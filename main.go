package main

import (
	"log"
	"net/http"

	"github.com/WalterPaes/BuscaCEP/cep"
)

func main() {

	http.HandleFunc("/cep/", cep.Handler)
	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
