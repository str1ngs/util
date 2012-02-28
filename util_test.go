package util

import (
	"testing"
)

var (
	existsFiles   = []string{"util.go", "util.go"}
	notExistFiles = []string{"aaaaaaaa", "bbbbbbbbb"}
)

func TestFileExists(t *testing.T) {
	for _, f := range existsFiles {
		exists := FileExists(f)
		if !exists {
			t.Errorf("expect to find %s got %v", f, exists)
		}
		t.Logf("%s -> %v", f, exists)
	}

	for _, f := range notExistFiles {
		exists := FileExists(f)
		if exists {
			t.Errorf("expect not to find %s got %v", f, exists)
		}
		t.Logf("%s -> %v", f, exists)
	}
}
