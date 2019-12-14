package viacep

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Search(cep string) ([]byte, error, int) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)

	res, err := http.Get(url)
	if err != nil {
		return nil, err, http.StatusInternalServerError
	}

	defer res.Body.Close()
	if res.StatusCode != 200  {
		return nil, errors.New("an error occurred"), res.StatusCode
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err, http.StatusInternalServerError
	}

	return body, nil, res.StatusCode
}