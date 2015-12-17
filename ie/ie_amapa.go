// ROTEIRO DE CRÍTICA DA INSCRIÇÃO ESTADUAL:
//  http://www.sintegra.gov.br/Cad_Estados/cad_AP.html

package ie

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
)

const (
	amapaIELenght      = 9
	amapaIEfirstDigits = "03"
)

// Amapa IE
type Amapa struct {
}

func (ieAmapa Amapa) assertValid(ie []int) (bool, error) {

	// validating the length
	if len(ie) != amapaIELenght {
		return false, errors.New(ieLenghtError)
	}

	// validating the first digits
	if strconv.Itoa(ie[0])+strconv.Itoa(ie[1]) != amapaIEfirstDigits {
		return false, fmt.Errorf(fmtfirstDigitsError, amapaIEfirstDigits)
	}

	// check digit
	p, d := getPandD(ie)

	newIE := make([]int, len(ie))
	copy(newIE, ie)
	newIE[len(newIE)-1] = p

	checkDigit := computeCheckDigit(newIE, getRulesAmapa(d))
	if checkDigit != ie[len(ie)-1] {
		return false, errors.New(invalidCheckDigits)
	}
	return true, nil
}

func (ieAmapa Amapa) generate() []int {
	ie := make([]int, alagoasIELenght)

	// fist digits
	ie[0] = 0
	ie[1] = 3

	// random numbers
	for i := 2; i < alagoasIELenght; i++ {
		ie[i] = rand.Intn(9)
	}

	p, d := getPandD(ie)

	newIE := make([]int, len(ie))
	copy(newIE, ie)
	newIE[len(newIE)-1] = p

	// check digits
	checkDigit := computeCheckDigit(newIE, getRulesAmapa(d))
	ie[len(ie)-1] = checkDigit
	return ie
}

func getPandD(ie []int) (int, int) {
	// Define-se dois valores, p e d, de acordo com as seguintes faixas de Inscrição Estadual:
	// De 03000001 a 03017000 => p = 5 e d = 0
	// De 03017001 a 03019022 => p = 9 e d = 1
	// De 03019023 em diante ===>p = 0 e d = 0
	var number string
	for i := 0; i < len(ie)-1; i++ {
		number += strconv.Itoa(ie[i])
	}

	value, _ := strconv.Atoi(number)
	p := 0
	d := 0
	switch {
	case value <= 3017000:
		p = 5
	case value >= 3017001 && value <= 3019022:
		p = 9
		d = 1
	}
	return p, d
}

func getRulesAmapa(d int) rules {
	return rules{
		initialMultiplier: 1,
		finalMultiplier:   9,
		substitute10:      0,
		substitute11:      d,
	}
}
