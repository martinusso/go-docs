package ie

import "testing"

var (
	validsIEAcre = [8]string{"01.004.823/001-12", "01.004.141/001-46",
		"01.001.349/001-77", "01.956.867/001-07", "01.379.333/036-16",
		"01.367.306/773-60", "01.658.566/892-98", "01.689.555/741-67"}
)

func Test_IEAcreLenght(t *testing.T) {
	ieAcre := Acre{}
	i, _ := assignStringToNumbers("123456789012")

	if _, err := ieAcre.assertValid(i); err == nil {
		t.Errorf("Unexpected success. Expected '%s'", ieLenghtError)
	}

	i, _ = assignStringToNumbers("12345678901234")
	if _, err := ieAcre.assertValid(i); err == nil {
		t.Errorf("Unexpected success. Expected '%s'", ieLenghtError)
	}
}

func Test_IEAcreFirstDigits(t *testing.T) {
	ieAcre := Acre{}

	i, _ := assignStringToNumbers("1234567890123")
	if _, err := ieAcre.assertValid(i); err == nil {
		t.Errorf("Unexpected success. Expected '%s'", firstDigitsError)
	}
}

func Test_IEAcreCheckDigits(t *testing.T) {
	ieAcre := Acre{}

	invalidIE, _ := assignStringToNumbers("0100482300122")
	if _, err := ieAcre.assertValid(invalidIE); err == nil {
		t.Errorf("Unexpected success. Expected '%s'", invalidCheckDigits)
	}
}

func Test_AssertValidWithIEAcre(t *testing.T) {
	for _, ie := range validsIEAcre {
		valid, err := AssertValid(ie, "AC")
		if !valid {
			t.Errorf("IE '%s' should be valid", ie)
		}
		if err != nil {
			t.Errorf("IE '%s' should be valid. Error: %s", ie, err.Error())
		}
	}
}

func Test_GenerateIEAcre(t *testing.T) {
	got, _ := Generate(ufAcre)
	if !Valid(got, ufAcre) {
		t.Errorf("IE %s is not valid", got)
	}
}
