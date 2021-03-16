package validator

import "testing"

func TestValidateURL(t *testing.T) {
	result := ValidateUrl("google.com")
	if !result {
		t.Errorf("Url not valid")
	}
}
