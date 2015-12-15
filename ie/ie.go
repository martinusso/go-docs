package ie

import (
	"errors"
	"strings"
)

const (
	invalidUF = "UF inválida."
)

// Valid validates the IE returning a boolean
func Valid(ie, uf string) bool {
	uf = strings.ToUpper(uf)
	isValid, err := AssertValid(ie, uf)
	if err != nil {
		return false
	}
	return isValid
}

// AssertValid validates the IE returning a boolean and the error if any
func AssertValid(ie, uf string) (bool, error) {
	if !validateUF(uf) {
		return false, errors.New(invalidUF)
	}
	ie = sanitize(ie)

	return false, nil
}

// Generate returns a random valid IE
func Generate(uf string) string {
	return ""
}

func sanitize(data string) string {
	data = strings.Replace(data, ".", "", -1)
	data = strings.Replace(data, "-", "", -1)
	data = strings.Replace(data, "/", "", -1)
	return data
}

func validateUF(uf string) bool {
	ufs := [27]string{"AC", "AL", "AM", "AP", "BA", "CE", "DF", "ES", "GO", "MA",
		"MG", "MS", "MT", "PA", "PB", "PE", "PI", "PR", "RJ", "RN", "RR", "RO", "RS",
		"SC", "SE", "SP", "TO"}

	for _, a := range ufs {
		if a == uf {
			return true
		}
	}
	return false
}
