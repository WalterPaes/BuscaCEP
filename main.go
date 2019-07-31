package main

import (
	"fmt"

	"github.com/WalterPaes/BuscaCEP/cep"
)

func main() {
	var c string

	fmt.Println("Insira o Cep que deseja buscar:")
	fmt.Scan(&c)

	cep.Search(c)
}
