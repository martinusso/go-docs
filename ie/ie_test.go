package ie

import "testing"

func Test_AssertValid(t *testing.T) {
	_, err := AssertValid("", "XX")
	if err == nil {
		t.Errorf("Unexpected success. Expected '%s'", invalidUF)
	}
}

func Test_AssignStringToIntArray(t *testing.T) {
	a, _ := assignStringToNumbers("123")
	if len(a) != 3 {
		t.Errorf("Invalid Lenght. Expected '%d' got '%d'", 3, len(a))
	}
	if a[0] != 1 {
		t.Errorf("Invalid element. Expected '%d' got '%d'", 1, a[0])
	}
	if a[1] != 2 {
		t.Errorf("Invalid element. Expected '%d' got '%d'", 2, a[1])
	}
	if a[2] != 3 {
		t.Errorf("Invalid element. Expected '%d' got '%d'", 3, a[2])
	}

	if _, err := assignStringToNumbers("a1"); err == nil {
		t.Errorf("Unexpected success. Expected parsing error")
	}
}
