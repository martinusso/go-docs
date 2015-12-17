// ROTEIRO DE CRÍTICA DA INSCRIÇÃO ESTADUAL:
//   http://www.sintegra.gov.br/Cad_Estados/cad_AC.html

package ie

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
)

const (
	ieAcreLenght      = 13
	acreIEfirstDigits = "01"
)

// Acre IE
type Acre struct {
}

func (ieAcre Acre) assertValid(ie []int) (bool, error) {
	// validating the length
	if len(ie) != ieAcreLenght {
		return false, errors.New(ieLenghtError)
	}
	// validating the first digits
	if strconv.Itoa(ie[0])+strconv.Itoa(ie[1]) != acreIEfirstDigits {
		return false, fmt.Errorf(fmtfirstDigitsError, acreIEfirstDigits)
	}

	checkDigit1 := computeCheckDigit(ie[:len(ie)-2], rulesDefault)
	checkDigit2 := computeCheckDigit(ie[:len(ie)-1], rulesDefault)
	if checkDigit1 != ie[len(ie)-2] || checkDigit2 != ie[len(ie)-1] {
		return false, errors.New(invalidCheckDigits)
	}
	return true, nil
}

func (ieAcre Acre) generate() []int {
	ie := make([]int, ieAcreLenght-2)

	// fist digits
	ie[0] = 0
	ie[1] = 1

	// random numbers
	for i := 2; i < ieAcreLenght-2; i++ {
		ie[i] = rand.Intn(9)
	}

	// check digits
	checkDigit1 := computeCheckDigit(ie, rulesDefault)
	ie = append(ie, checkDigit1)
	checkDigit2 := computeCheckDigit(ie, rulesDefault)
	ie = append(ie, checkDigit2)

	return ie
}
