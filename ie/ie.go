package ie

import (
	"errors"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const (
	invalidUF           = "UF inválida."
	ieLenghtError       = "Tamanho da IE inválido."
	invalidCheckDigits  = "Dígito Verificador inválido"
	firstDigitsError    = "Incorrect first digits"
	fmtfirstDigitsError = "Os primeiros dois dígitos são sempre %s"

	ufAcre    = "AC"
	ufAlagoas = "AL"
	ufAmapa   = "AP"
)

// IE interface to validation and generation of IE
type IE interface {
	assertValid(ie []int) (bool, error)
	generate() []int
}

// Valid validates the IE returning a boolean
func Valid(ie, uf string) bool {
	isValid, err := AssertValid(ie, uf)
	if err != nil {
		return false
	}
	return isValid
}

// AssertValid validates the IE returning a boolean and the error if any
func AssertValid(ie, uf string) (bool, error) {
	uf = strings.ToUpper(uf)
	ie = sanitize(ie)

	numbers, err := assignStringToNumbers(ie)
	if err != nil {
		return false, err
	}

	switch uf {
	case ufAcre:
		return Acre{}.assertValid(numbers)
	case ufAlagoas:
		return Alagoas{}.assertValid(numbers)
	case ufAmapa:
		return Amapa{}.assertValid(numbers)
	default:
		return false, errors.New(invalidUF)
	}
}

// Generate returns a random valid IE
func Generate(uf string) (string, error) {
	rand.Seed(time.Now().UTC().UnixNano())

	uf = strings.ToUpper(uf)

	var numbers []int
	switch uf {
	case ufAcre:
		numbers = Acre{}.generate()
	case ufAlagoas:
		numbers = Alagoas{}.generate()
	default:
		return "", errors.New(invalidUF)
	}

	var str string
	for _, value := range numbers {
		str += strconv.Itoa(value)
	}
	return str, nil
}

func sanitize(data string) string {
	data = strings.Replace(data, ".", "", -1)
	data = strings.Replace(data, "-", "", -1)
	data = strings.Replace(data, "/", "", -1)
	return data
}

func assignStringToNumbers(data string) ([]int, error) {
	a := make([]int, len(data))
	for i, s := range strings.Split(data, "") {
		original, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}

		a[i] = original
	}
	return a, nil
}

func computeCheckDigit(data []int, rules rules) int {
	lenghtMultipliers := rules.finalMultiplier - rules.initialMultiplier + 1
	multipliers := make([]int, lenghtMultipliers)
	for i, m := 0, rules.initialMultiplier; m <= rules.finalMultiplier; i++ {
		multipliers[i] = m
		m++
	}
	modulus := 11
	sum := 0

	for i, m := len(data)-1, 0; i >= 0; i-- {
		sum += data[i] * multipliers[m]

		m++
		if m >= len(multipliers) {
			m = 0
		}
	}

	mod := int(math.Mod(float64(sum), 11))
	r := modulus - mod
	// fmt.Println("10", rules.substitute10)
	// fmt.Println("11", rules.substitute11)
	switch r {
	case 10:
		return rules.substitute10
	case 11:
		return rules.substitute11
	default:
		return r
	}
}
