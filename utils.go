package utils

import (
	"fmt"
	"os"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

// 추후에는 내용이 많아지면 분리하자.
// tag 추가 중.
var (
	pTrue = true
	PTrue = &pTrue

	pFalse = false
	PFalse = &pFalse
)

// IsEmptyString returns true if the string is empty or contains only whitespace, false otherwise
func IsEmptyString(s string) bool {
	return len(strings.TrimSpace(s)) == 0
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

// FileExists true if the file exists, false if the file does not exist
// If the file exists, the FileInfo of the file is returned.
func FileExists(path string) (bool, os.FileInfo, error) {
	var (
		fileInfo os.FileInfo
		err      error
	)
	if IsEmptyString(path) {
		return false, nil, fmt.Errorf("path is emtpy")
	}
	if fileInfo, err = os.Stat(path); os.IsNotExist(err) {
		return false, nil, nil
	}
	return true, fileInfo, nil
}

// Truncate
func Truncate(path string) error {

	file, err := os.OpenFile(path, os.O_RDWR, 0666)
	defer file.Close()

	if err != nil {
		return err
	}

	err = file.Truncate(0)
	if err != nil {
		return err
	}
	_, err = file.Seek(0, 0)
	if err != nil {
		return err
	}

	return nil
}

// TODO 따로 빼놓자.
// https://stackoverflow.com/questions/37334119/how-to-delete-an-element-from-a-slice-in-golang
// https://yourbasic.org/golang/delete-element-slice/
func Remove(ss []chan interface{}, i int) []chan interface{} {

	copy(ss[i:], ss[i+1:]) // Shift a[i+1:] left one index.
	ss[len(ss)-1] = nil    // Erase last element (write zero value).
	ss = ss[:len(ss)-1]    // Truncate slice.

	return ss
	//return append(ss[:i], ss[i+1:]...)
}

// Contains checks if a slice contains a given string
func Contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
