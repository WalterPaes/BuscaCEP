package tests

import (
	"bytes"
	"github.com/WalterPaes/BuscaCEP/pkg/domains/cep"
	"io/ioutil"
	"net/http"
	"testing"
)

type ServiceCepMock struct {}

func (svc ServiceCepMock) Search(_ string) ([]byte, error, int) {
	cepJson := `{"cep": "01001-000", "logradouro": "Praça da Sé", "complemento": "lado ímpar",
			"bairro": "Sé", "localidade": "São Paulo", "uf": "SP", "unidade": "",
			"ibge": "3550308", "gia": "1004"}`
	responseBody, _ := ioutil.ReadAll(bytes.NewBufferString(cepJson))
	return responseBody, nil, http.StatusOK
}

func TestSearchCep(t *testing.T) {
	address := cep.NewAddress("01001000", &ServiceCepMock{})
	result, err, status := address.SearchCep()
	responseBody := `{"cep":"01001-000","logradouro":"Praça da Sé","complemento":"lado ímpar","bairro":"Sé","localidade":"São Paulo","uf":"SP","ibge":"3550308","gia":"1004"}`

	if string(result) != string(responseBody) {
		t.Errorf("The Response should be: '%s', Expected: '%s'", string(result), string(responseBody))
	}

	if err != nil {
		t.Errorf("The Error should be: '%v', Expected: '%v'", err, nil)
	}

	if status != http.StatusOK {
		t.Errorf("The Http Status Code should be: '%v', Expected: '%v'", err, nil)
	}
}