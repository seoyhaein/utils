package utils

import (
	"strings"
)

var (
	pTrue = true
	PTrue = &pTrue

	pFalse = false
	PFalse = &pFalse
)

//IsEmptyString true if string is empty, false otherwise
func IsEmptyString(s string) bool {

	r := len(strings.TrimSpace(s))

	if r == 0 {
		return true
	}
	return false
}
