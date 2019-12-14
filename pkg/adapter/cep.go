package adapter

import (
	"github.com/WalterPaes/BuscaCEP/pkg/domains/cep"
	"github.com/WalterPaes/BuscaCEP/pkg/helper"
	"net/http"
)

func CepSearch(c string) ([]byte, int) {

	// Checking if is a valid PostalCode
	err := cep.Validation(c)
	if err != nil {
		var e helper.Error
		e.Message = err.Error()
		return e.ToByte(), http.StatusBadRequest
	}

	// Make a Search by Postal Code Number
	result, err, status := cep.Search(c)
	if err != nil {
		var e helper.Error
		e.Message = err.Error()
		return e.ToByte(), status
	}

	return result, status
}
