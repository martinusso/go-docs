package ie

import "testing"

func Test_AssertValid(t *testing.T) {
	_, err := AssertValid("", "XX")
	if err == nil {
		t.Errorf("Unexpected success. Expected '%s'", invalidUF)
	}
}
