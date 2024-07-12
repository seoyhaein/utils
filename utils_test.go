package utils

import (
	"os"
	"testing"
)

// TestIsEmptyString tests the IsEmptyString function
func TestIsEmptyString(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"", true},
		{" ", true},
		{"a", false},
		{" test ", false},
	}

	for _, test := range tests {
		result := IsEmptyString(test.input)
		if result != test.expected {
			t.Errorf("IsEmptyString(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

// TestDeepCopy tests the DeepCopy function
func TestDeepCopy(t *testing.T) {
	type TestStruct struct {
		A int
		B string
	}

	src := TestStruct{A: 1, B: "test"}
	var dst TestStruct

	err := DeepCopy(&dst, src)
	if err != nil {
		t.Errorf("DeepCopy failed: %v", err)
	}

	if dst != src {
		t.Errorf("DeepCopy result = %v; want %v", dst, src)
	}
}

// TestFileExists tests the FileExists function
func TestFileExists(t *testing.T) {
	// Test with a non-existing file
	exists, _, err := FileExists("non_existent_file.txt")
	if err != nil {
		t.Errorf("FileExists failed: %v", err)
	}
	if exists {
		t.Errorf("FileExists(\"non_existent_file.txt\") = %v; want %v", exists, false)
	}

	// Test with an existing file
	testFileName := "test_file.txt"
	_, err = os.Create(testFileName)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(testFileName)

	exists, _, err = FileExists(testFileName)
	if err != nil {
		t.Errorf("FileExists failed: %v", err)
	}
	if !exists {
		t.Errorf("FileExists(%q) = %v; want %v", testFileName, exists, true)
	}
}

// TestTruncate tests the Truncate function
func TestTruncate(t *testing.T) {
	testFileName := "test_file.txt"
	file, err := os.Create(testFileName)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	file.WriteString("some content")
	file.Close()
	defer os.Remove(testFileName)

	err = Truncate(testFileName)
	if err != nil {
		t.Errorf("Truncate failed: %v", err)
	}

	stat, err := os.Stat(testFileName)
	if err != nil {
		t.Errorf("Failed to stat test file: %v", err)
	}

	if stat.Size() != 0 {
		t.Errorf("File size after Truncate = %d; want %d", stat.Size(), 0)
	}
}

// TestRemove tests the Remove function
func TestRemove(t *testing.T) {
	ch1 := make(chan interface{})
	ch2 := make(chan interface{})
	ch3 := make(chan interface{})
	ss := []chan interface{}{ch1, ch2, ch3}

	ss = Remove(ss, 1)

	if len(ss) != 2 {
		t.Errorf("Remove length = %d; want %d", len(ss), 2)
	}
	if ss[0] != ch1 || ss[1] != ch3 {
		t.Errorf("Remove result = %v; want [%v %v]", ss, ch1, ch3)
	}
}
