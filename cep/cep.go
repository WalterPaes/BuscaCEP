package cep

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Cep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
}

func Search(cep string) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	var c Cep
	json.Unmarshal([]byte(body), &c)

	c.Show()
}

func (c Cep) Show() {
	fmt.Printf("CEP: %s\n", c.Cep)
	fmt.Printf("Logradouro: %s\n", c.Logradouro)
	fmt.Printf("Complemento: %s\n", c.Complemento)
	fmt.Printf("Bairro: %s\n", c.Bairro)
	fmt.Printf("Localidade: %s\n", c.Localidade)
	fmt.Printf("UF: %s\n", c.Uf)
}
