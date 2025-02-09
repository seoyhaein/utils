package utils

import (
	"fmt"
	"os"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

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

// CheckPath validates and normalizes the provided file path.
// It returns an error if the file path is empty, otherwise it returns the cleaned file path.
func CheckPath(filePath string) (string, error) {
	// Check if the file path is empty using a helper function (e.g., IsEmptyString)
	if IsEmptyString(filePath) {
		return "", fmt.Errorf("file path cannot be empty")
	}
	// filepath.Clean standardizes the path by removing redundant separators,
	// resolving dot (".") elements, and simplifying relative path components.
	return filepath.Clean(filePath), nil
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
	if IsEmptyString(path) {
		return false, nil, fmt.Errorf("path is empty")
	}

	fileInfo, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil, nil
		}
		return false, nil, err
	}
	return true, fileInfo, nil
}

// Truncate 일단 수정했음. 24/11/15 by seoyhaein
func Truncate(path string) error {

	file, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer func() {
		if cerr := file.Close(); cerr != nil {
			fmt.Printf("Error closing file: %v\n", cerr)
		}
	}()

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
