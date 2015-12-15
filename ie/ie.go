package ie

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

const (
	invalidUF          = "UF inválida."
	ieLenghtError      = "Tamanho da IE inválido."
	invalidCheckDigits = "Dígito Verificador inválido"
)

// IE interface to validation and generation of IE
type IE interface {
	AssertValid(ie []int) (bool, error)
	Generate(uf string) string
}

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
	numbers, err := assignStringToNumbers(ie)
	if err != nil {
		return false, err
	}
	return Acre{}.AssertValid(numbers)
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

/*
private string GetDigitSum()
        {
            int soma = 0;
            for (int i = this.numero.Length - 1, m = 0; i >= 0; i--)
            {
                int produto = (int)char.GetNumericValue(this.numero[i]) * this.multiplicadores[m];
                soma += somarAlgarismos ? SomaAlgarismos(produto) : produto;

                if (++m >= this.multiplicadores.Count) m = 0;
            }

            int mod = (soma % modulo);
            int resultado = complementarDoModulo ? modulo - mod : mod;

            if (substituicoes.ContainsKey(resultado))
            {
                return this.substituicoes[resultado];
            }

            return resultado.ToString();
        }
*/
