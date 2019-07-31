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
	var c Cep
	cep := strings.TrimPrefix(request.URL.Path, "/cep/")

	switch {
	case request.Method == "GET" && cep != "":
		response := c.Search(cep)
		writer.Header().Set("Content-Type", "application/json")
		fmt.Fprint(writer, string(response))
	default:
		writer.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(writer, "NOT FOUND")
	}
}

func (c Cep) Search(cep string) []byte {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	json.Unmarshal([]byte(body), &c)

	json, _ := json.Marshal(c)

	return json
}
