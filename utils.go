package utils

import (
	"fmt"
	"os"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

// 추후에는 내용이 많아지면 분리하자.
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

func FileExists(path string) (bool, error) {
	if IsEmptyString(path) {
		return false, fmt.Errorf("path is emtpy")
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false, fmt.Errorf("the file does not exist")
	}

	return true, nil
}
