package viacep

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

var ErrOccurred = errors.New("an error occurred")

type Service struct {}

// Search a Postal Code in ViaCEP service
func (svc Service) Search(cep string) ([]byte, error, int) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)

	// Do Request
	res, err := http.Get(url)
	if err != nil {
		return nil, err, http.StatusInternalServerError
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, ErrOccurred, res.StatusCode
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err, http.StatusInternalServerError
	}

	return body, nil, res.StatusCode
}
