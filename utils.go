package utils

import (
	"fmt"
	"os"
	"path/filepath"
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
// 문자열이 공백만 있거나 비어 있으면 true, 아니면 false 반환
func IsEmptyString(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

// CheckPath validates and normalizes the provided file path.
// It returns an error if the file path is empty, otherwise it returns the cleaned file path.
// 파일 경로를 검사하고 정규화함. 비어 있으면 에러, 아니면 정리된 경로 반환
func CheckPath(filePath string) (string, error) {
	// Check if the file path is empty using a helper function (e.g., IsEmptyString)
	// 파일 경로가 비어있는지 검사함.
	if IsEmptyString(filePath) {
		return "", fmt.Errorf("file path cannot be empty")
	}
	// filepath.Clean standardizes the path by removing redundant separators,
	// resolving dot '.' elements, and simplifying relative path components.
	// 경로의 불필요한 구분자 및 상대경로 등을 정리함.
	return filepath.Clean(filePath), nil
}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// DeepCopy does a deep copy of a structure
// Error checking of parameters delegated to json engine
// JSON 직렬화/역직렬화를 이용해 구조체의 깊은 복사를 수행함.
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
// 주어진 경로에 파일이 존재하는지 확인함. 있으면 true 와 FileInfo 반환, 없으면 false 반환
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

// Truncate clears the contents of a file by truncating it to zero length.
// 파일 내용을 0 길이로 잘라서 삭제함.
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

// Remove deletes an element from a slice of channels at the specified index (preserving order).
// 슬라이스에서 특정 인덱스의 채널 요소를 제거함 (순서 유지)
func Remove(ss []chan interface{}, i int) []chan interface{} {

	copy(ss[i:], ss[i+1:]) // Shift a[i+1:] left one index.
	ss[len(ss)-1] = nil    // Erase last element (write zero value).
	ss = ss[:len(ss)-1]    // Truncate slice.

	return ss
	//return append(ss[:i], ss[i+1:]...)
}

// Contains checks if a slice contains a given string
// 슬라이스에 주어진 문자열이 포함되어 있는지 확인함
func Contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// ExcludeFiles returns true if fileName exactly matches any string in exclusions.
// 파일 이름이 exclusions 목록에 정확히 일치하는지 검사함
func ExcludeFiles(fileName string, exclusions []string) bool {
	for _, ex := range exclusions {
		if fileName == ex {
			return true
		}
	}
	return false
}
