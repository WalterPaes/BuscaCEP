package adapter

import (
	"github.com/WalterPaes/BuscaCEP/pkg/domains/cep"
	"github.com/WalterPaes/BuscaCEP/pkg/helper"
	"github.com/WalterPaes/BuscaCEP/pkg/integration/viacep"
)

func SearchCep(c string) ([]byte, int) {
	address := cep.NewAddress(c, &viacep.Service{})

	// Make a Search by Zip Code Number
	result, err, status := address.SearchCep()
	if err != nil {
		var e helper.Error
		e.Message = err.Error()
		return e.ToByte(), status
	}

	return result, status
}
