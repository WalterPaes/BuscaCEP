package cep

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/WalterPaes/BuscaCEP/pkg/integration/viacep"
	"net/http"
	"regexp"
)

func Validation(cep string) error {

	reg, _ := regexp.Compile("[^0-9]+")
	cep = reg.ReplaceAllString(cep, "")
	fmt.Println("a", cep)
	if len(cep) != 8 {
		return errors.New("invalid postal code")
	}
	return nil
}

func Search(cep string) ([]byte, error, int){
	result, err, status := viacep.Search(cep)

	var data Data
	err = json.Unmarshal(result, &data)
	if err != nil {
		return nil, err, http.StatusBadRequest
	}

	if data.Cep == "" {
		return nil, errors.New("not found"), http.StatusNotFound
	}

	return result, err, status
}