package cep

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Cep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
}

func Handler(writer http.ResponseWriter, request *http.Request) {
	var cep Cep

	cep.Cep = strings.TrimPrefix(request.URL.Path, "/cep/")
	if len(cep.Cep) != 8 {
		writer.Header().Set("Content-Type", "application/json")
		fmt.Fprint(writer, string("{\"erro\":\"Cep Inválido\"}"))
		return
	}

	switch {
	case request.Method == "GET":
		response, err := cep.Search()

		if len(err) > 0 {
			writer.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(writer, string("{\"erro\":\""+err+"\"}"))
		} else {
			writer.Header().Set("Content-Type", "application/json")
			fmt.Fprint(writer, response)
		}
	default:
		writer.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(writer, "CEP Inválido!")
	}
	return
}

func (c *Cep) Search() (string, string) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", c.Cep)

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(body, &c)

	if len(c.Uf) < 1 {
		return "", "CEP Não existente!"
	} else {
		json, _ := json.Marshal(c)
		return string(json), ""
	}
}
