package cpf

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

	for i := 0; i <= 9; i++ {
		invalidCPF := strings.Repeat(string(i), 11)

		isValid, err = AssertValid(invalidCPF)
		if err == nil {
			t.Errorf("Shouldn't be nil")
		}
		if isValid {
			t.Errorf("CPF can't be consist by repeated numbers.")
		}
		if err.Error() != repeatedDigits {
			t.Errorf("Expected %s got %s", repeatedDigits, err.Error())
		}
	}
	isValid, err = AssertValid("00000000000000")
}

func Test_Valid(t *testing.T) {
	validCPF := "52998224725"
	formattedCPF := "529.982.247-25"
	invalidCPF := "52295224717"

	if !Valid(validCPF) {
		t.Errorf("CPF should be valid")
	}

	if !Valid(formattedCPF) {
		t.Errorf("Formatted CPF should be valid")
	}

	if Valid(invalidCPF) {
		t.Errorf("CPF Shouldn't be valid")
	}

	for i := 0; i <= 9; i++ {
		invalidCPF := strings.Repeat(string(i), 11)
		if Valid(invalidCPF) {
			t.Errorf("CPF can't be consist by repeated numbers.")
		}
	}
}
