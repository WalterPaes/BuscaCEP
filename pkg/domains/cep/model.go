package cep

import (
	"encoding/json"
	"errors"
	"net/http"
	"regexp"
)

var (
	ErrInvalidZipCode = errors.New("invalid postal code")
	ErrNotFound       = errors.New("not found")
)

type ServiceCep interface {
	Search(cep string) ([]byte, error, int)
}

type Address struct {
	CEP         string     `json:"cep"`
	Logradouro  string     `json:"logradouro"`
	Complemento string     `json:"complemento"`
	Bairro      string     `json:"bairro"`
	Localidade  string     `json:"localidade"`
	UF          string     `json:"uf"`
	Unidade     string     `json:"unidade,omitempty"`
	IBGE        string     `json:"ibge,omitempty"`
	Gia         string     `json:"gia,omitempty"`
	svc         ServiceCep `json:"-"`
}

func NewAddress(cep string, service ServiceCep) *Address {
	return &Address{CEP: cep, svc: service}
}

// Search By CEP
func (ad *Address) SearchCep() ([]byte, error, int) {
	err := validation(ad.CEP)
	if err != nil {
		return nil, err, http.StatusBadRequest
	}

	result, err, status := ad.svc.Search(ad.CEP)
	if err != nil {
		return result, err, status
	}

	err = json.Unmarshal(result, &ad)
	if err != nil {
		return nil, err, http.StatusInternalServerError
	}

	if ad.Logradouro == "" {
		return nil, ErrNotFound, http.StatusNotFound
	}

	result, err = json.Marshal(ad)
	return result, err, status
}

func validation(cep string) error {
	reg, _ := regexp.Compile("[^0-9]+")
	cep = reg.ReplaceAllString(cep, "")
	if len(cep) != 8 {
		return ErrInvalidZipCode
	}
	return nil
}
