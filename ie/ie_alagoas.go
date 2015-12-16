// ROTEIRO DE CRÍTICA DA INSCRIÇÃO ESTADUAL:
//  http://www.sintegra.gov.br/Cad_Estados/cad_AL.html

package ie

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
)

const (
	alagoasIELenght      = 9
	alagoasIEfirstDigits = "24"
	invalidCompanyType   = "Tipo de empresa inválido."
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
		return false, errors.New(invalidCompanyType)
	}

	checkDigit := computeCheckDigit(ie[:len(ie)-1])
	if checkDigit != ie[len(ie)-1] {
		return false, errors.New(invalidCheckDigits)
	}
	return true, nil
}

func (ieAlagoas Alagoas) generate() []int {
	ie := make([]int, alagoasIELenght-1)

	// fist digits
	ie[0] = 2
	ie[1] = 4

	// company type
	ie[2] = companyType[rand.Intn(5)]

	// random numbers
	for i := 3; i < alagoasIELenght-1; i++ {
		ie[i] = rand.Intn(9)
	}

	// check digits
	checkDigit := computeCheckDigit(ie)

	return append(ie, checkDigit)
}
