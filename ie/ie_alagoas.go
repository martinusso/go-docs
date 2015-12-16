// ROTEIRO DE CRÍTICA DA INSCRIÇÃO ESTADUAL:
//  http://www.sintegra.gov.br/Cad_Estados/cad_AL.html

package ie

import (
	"errors"
	"fmt"
	"strconv"
)

const (
	alagoasIEfirstDigits = "24"
	companyTypeError     = "Tipo de empresa inválido."
)

var (
	// 0-Normal, 3-Produtor Rural, 5-Substituta, 7- Micro-Empresa Ambulante, 8-Micro-Empresa
	companyType = [5]int{0, 3, 5, 7, 8}
)

// Alagoas IE
type Alagoas struct {
}

func (ieAlagoas Alagoas) assertValid(ie []int) (bool, error) {
	// validating the first digits
	if strconv.Itoa(ie[0])+strconv.Itoa(ie[1]) != alagoasIEfirstDigits {
		return false, fmt.Errorf(fmtfirstDigitsError, alagoasIEfirstDigits)
	}

	// validating company type
	companyTypeIsValid := false
	for _, t := range companyType {
		if t == ie[2] {
			companyTypeIsValid = true
		}
	}
	if !companyTypeIsValid {
		return false, fmt.Errorf(fmtfirstDigitsError, alagoasIEfirstDigits)
	}

	checkDigit1 := computeCheckDigit(ie[:len(ie)-1])
	if checkDigit1 != ie[len(ie)-1] {
		return false, errors.New(invalidCheckDigits)
	}
	return true, nil
}
