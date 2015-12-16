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
)

// IE interface to validation and generation of IE
type IE interface {
	AssertValid(ie []int) (bool, error)
	Generate() string
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
		return Acre{}.AssertValid(numbers)
	case ufAlagoas:
		return Alagoas{}.assertValid(numbers)
	default:
		return false, errors.New(invalidUF)
	}
}

// Generate returns a random valid IE
func Generate(uf string) (string, error) {
	rand.Seed(time.Now().UTC().UnixNano())

	uf = strings.ToUpper(uf)
	if !validateUF(uf) {
		return "", errors.New(invalidUF)
	}

	numbers := Acre{}.Generate()

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

func computeCheckDigit(data []int) int {
	multipliers := [8]int{2, 3, 4, 5, 6, 7, 8, 9}
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

	if r > 9 {
		return 0
	}
	return r
}
