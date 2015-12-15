// ROTEIRO DE CRÍTICA DA INSCRIÇÃO ESTADUAL:
//   http://www.sintegra.gov.br/Cad_Estados/cad_AC.html

package ie

import (
	"errors"
	"strconv"
)

const (
	ieAcreLenght     = 13
	firstDigits      = "01"
	firstDigitsError = "Os primeiros dois dígitos são sempre 01"
)

// Acre IE
type Acre struct {
}

// AssertValid validates the Acre IE returning a boolean and the error if any
func (ieAcre Acre) AssertValid(ie []int) (bool, error) {
	// validating the length
	if len(ie) != ieAcreLenght {
		return false, errors.New(ieLenghtError)
	}
	// validating the first digits
	if strconv.Itoa(ie[0])+strconv.Itoa(ie[1]) != firstDigits {
		return false, errors.New(firstDigitsError)
	}

	checkDigit1 := computeCheckDigit(ie[:len(ie)-2])
	checkDigit2 := computeCheckDigit(ie[:len(ie)-1])
	if checkDigit1 != ie[len(ie)-2] || checkDigit2 != ie[len(ie)-1] {
		return false, errors.New(invalidCheckDigits)
	}
	return true, nil
}
