package ie

import "testing"

var (
	validsAmapaIE = [3]string{"03.012.345-9", "03.002.547-3", "030181234"}
)

func Test_AmapaIELenght(t *testing.T) {
	ieAmapa := Amapa{}
	i, _ := assignStringToNumbers("03012345")

	if _, err := ieAmapa.assertValid(i); err == nil {
		t.Errorf("Unexpected success. Expected '%s'", ieLenghtError)
	}

	i, _ = assignStringToNumbers("0301234590")
	if _, err := ieAmapa.assertValid(i); err == nil {
		t.Errorf("Unexpected success. Expected '%s'", ieLenghtError)
	}
}

func Test_AmapaIEFirstDigits(t *testing.T) {
	ieAmapa := Amapa{}

	i, _ := assignStringToNumbers("040123459")
	if _, err := ieAmapa.assertValid(i); err == nil {
		t.Errorf("Unexpected success. Expected '%s'", firstDigitsError)
	}
}

func Test_AmapaIECheckDigits(t *testing.T) {
	ieAmapa := Amapa{}

	invalidIE, _ := assignStringToNumbers("030123451")
	if _, err := ieAmapa.assertValid(invalidIE); err == nil {
		t.Errorf("Unexpected success. Expected '%s'", invalidCheckDigits)
	}
}

func Test_AssertValidWithAmapaIE(t *testing.T) {
	for _, ie := range validsAmapaIE {
		valid, err := AssertValid(ie, ufAmapa)
		if !valid {
			t.Errorf("IE '%s' should be valid", ie)
		}
		if err != nil {
			t.Errorf("IE '%s' should be valid. Error: %s", ie, err.Error())
		}
	}
}

func Test_GenerateAmapaIE(t *testing.T) {
	got, _ := Generate(ufAmapa)
	if !Valid(got, ufAmapa) {
		t.Errorf("IE %s is not valid", got)
	}
}
