package cpf

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

	for i := 0; i <= 9; i++ {
		invalidCPF := strings.Repeat(strconv.Itoa(i), 11)

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

func TestValid(t *testing.T) {
	listOfValidCPF := []string{"52998224725", "529.982.247-25"}
	for _, cpf := range listOfValidCPF {
		if !Valid(cpf) {
			t.Errorf("Test Failed. CPF %s must be valid.", cpf)
		}
	}

	if Valid("12E45678901") {
		t.Errorf("Test Failed. Only numbers are allowed.")
	}

	if Valid("52295224717") {
		t.Errorf("CPF Shouldn't be valid")
	}

	for i := 0; i <= 9; i++ {
		invalidCPF := strings.Repeat(strconv.Itoa(i), 11)
		if Valid(invalidCPF) {
			t.Errorf("CPF can't be consist by repeated numbers.")
		}
	}
}

func TestGenerate(t *testing.T) {
	got := Generate()
	if !Valid(got) {
		t.Errorf("Generated CPF should be valid")
	}
}
