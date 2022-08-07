package utils

import (
	"strings"

	jsoniter "github.com/json-iterator/go"
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

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// DeepCopy does a deep copy of a structure
// Error checking of parameters delegated to json engine
var DeepCopy = func(dst interface{}, src interface{}) error {
	payload, err := json.Marshal(src)
	if err != nil {
		return err
	}

	err = json.Unmarshal(payload, dst)
	if err != nil {
		return err
	}
	return nil
}
