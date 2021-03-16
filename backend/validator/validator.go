package validator

import (
	"net/url"
	"strings"
)

func ValidateUrl(s string) bool {
	if s == "" || strings.HasPrefix(s, ".") {
		return false
	}

	temp := s
	if strings.Contains(s, ":") && !strings.Contains(s, "://") {
		temp = "http://" + s
	}

	u, err := url.Parse(temp)
	if err != nil {
		return false
	}

	if strings.HasPrefix(u.Host, ".") {
		return false
	}

	if u.Host == "" && (u.Path != "" && !strings.Contains(u.Path, ".")) {
		return false
	}

	return true
}
