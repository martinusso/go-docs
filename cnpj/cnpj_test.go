package cnpj

import (
	"strconv"
	"strings"
	"testing"
)

func TestAssertValid(t *testing.T) {
	isValid, err := AssertValid("123456789012345")
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

func TestValid(t *testing.T) {
	if Valid("1234567890") {
		t.Errorf("Test Failed. CNPJ length must de invalid")
	}
	if Valid("99999999000100") {
		t.Errorf("Test Failed. CNPJ %s must be invalid.", "99999999000100")
	}

	listOfValidCNPJ := []string{"99999999000191", "99.999.999/0001-91", "20.717.607/0001-02"}
	for _, cnpj := range listOfValidCNPJ {
		if !Valid(cnpj) {
			t.Errorf("Test Failed. CNPJ %s must be valid.", cnpj)
		}
	}

	for i := 0; i <= 9; i++ {
		invalidCNPJ := strings.Repeat(strconv.Itoa(i), 14)
		if Valid(invalidCNPJ) {
			t.Errorf("CNPJ is invalid (%s)", invalidCNPJ)
		}
	}
}

func TestGenerate(t *testing.T) {
	got := Generate()
	if !Valid(got) {
		t.Errorf("CNPJ is not valid")
	}
}
