package validator

import "testing"

func TestValidateURLPass(t *testing.T) {
	result := ValidateUrl("google.com")
	if !result {
		t.Errorf("Url not valid")
	}
}
