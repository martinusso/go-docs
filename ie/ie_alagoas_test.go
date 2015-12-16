package ie

import "testing"

var (
	validsIEAlagoas = [16]string{
		"24.076.739-0", "24.089.826-5", "24.099.991-6", "24.067.173-2",
		"24.079.990-9", "24.089.451-0", "24.080.152-0", "24.092.497-5",
		"24.088.932-0", "24.097.262-7", "24.086.162-0", "24.097.871-4",
		"24.085.016-5", "24.073.874-8", "24.071.760-0", "24.065.706-3"}
)

func Test_AlagoasIEFirstDigits(t *testing.T) {
	alagoasIE := Alagoas{}

	i, _ := assignStringToNumbers("120456789")
	if _, err := alagoasIE.assertValid(i); err == nil {
		t.Errorf("Unexpected success. Expected '%s'", firstDigitsError)
	}
}

func Test_AlagoasIECompanyType(t *testing.T) {
	alagoasIE := Alagoas{}

	listInvalidCompanyType := [5]int{1, 2, 4, 6, 9}

	for _, i := range listInvalidCompanyType {
		ieModel := []int{2, 4, i, 1, 2, 3, 4, 5, 6}
		if _, err := alagoasIE.assertValid(ieModel); err == nil {
			t.Errorf("Unexpected success. Expected invalid company type. %v %d", ieModel, i)
		}
	}
}

func Test_AlagoasIECheckDigits(t *testing.T) {
	alagoasIE := Alagoas{}

	invalidIE, _ := assignStringToNumbers("240767391")
	if _, err := alagoasIE.assertValid(invalidIE); err == nil {
		t.Errorf("Unexpected success. Expected '%s'", invalidCheckDigits)
	}
}

func Test_AssertValidWithAlagoasIE(t *testing.T) {
	for _, ie := range validsIEAlagoas {
		valid, err := AssertValid(ie, ufAlagoas)
		if !valid {
			t.Errorf("IE '%s' should be valid", ie)
		}
		if err != nil {
			t.Errorf("IE '%s' should be valid. Error: %s", ie, err.Error())
		}
	}
}

func Test_GenerateAlagoasIE(t *testing.T) {
	got, _ := Generate(ufAlagoas)
	if !Valid(got, ufAlagoas) {
		t.Errorf("IE %s is not valid", got)
	}
}
