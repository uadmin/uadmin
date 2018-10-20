package helper

import (
	"regexp"
)

// ValidateSubdomain !
func ValidateSubdomain(email string) bool {
	pattern := regexp.MustCompile("[a-z\\-0-9]{3,24}")
	return pattern.MatchString(email)
}
