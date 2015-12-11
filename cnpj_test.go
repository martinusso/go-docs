package cnpj

import (
	"strings"
	"testing"
)

func Test_AssertValid(t *testing.T) {
	invalidSize := "123456789012345"
	isValid, err := AssertValid(invalidSize)
	if err == nil {
		t.Errorf("Shouldn't be nil")
	}
	if isValid {
		t.Errorf("Shouldn't be true")
	}
	if err.Error() != invalidLength {
		t.Errorf("Expected %s got %s", invalidLength, err.Error())
	}

	isValid, err = AssertValid("00000000000000")
	if err == nil {
		t.Errorf("Shouldn't be nil")
	}
	if isValid {
		t.Errorf("Shouldn't be true")
	}
	if err.Error() != repeatedDigits {
		t.Errorf("Expected %s got %s", repeatedDigits, err.Error())
	}
}

func Test_Valid(t *testing.T) {
	validCNPJ := "99999999000191"
	formattedCNPJ := "99.999.999/0001-91"
	invalidCNPJ := "99999999000100"

	if got := Valid(validCNPJ); got != true {
		t.Errorf("CNPJ is not valid")
	}

	if got := Valid(formattedCNPJ); got != true {
		t.Errorf("CNPJ is not valid")
	}

	if got := Valid(invalidCNPJ); got != false {
		t.Errorf("CNPJ is valid")
	}

	for i := 0; i <= 9; i++ {
		invalidCNPJ := strings.Repeat(string(i), 14)

		if got := Valid(invalidCNPJ); got != false {

			t.Errorf("CNPJ is valid")
		}
	}
}

func Test_Generate(t *testing.T) {
	got := Generate()
	if !Valid(got) {
		t.Errorf("CNPJ is not valid")
	}
}
