package file

import (
	"os"
	"path"
	"testing"
)

var (
	existsFiles   = []string{"../util.go", "file.go"}
	notExistFiles = []string{"aaaaaaaa", "bbbbbbbbb"}
)

func TestExists(t *testing.T) {
	for _, f := range existsFiles {
		exists := Exists(f)
		if !exists {
			t.Errorf("expect to find %s got %v", f, exists)
		}
		t.Logf("%s -> %v", f, exists)
	}
	for _, f := range notExistFiles {
		exists := Exists(f)
		if exists {
			t.Errorf("expect not to find %s got %v", f, exists)
		}
		t.Logf("%s -> %v", f, exists)
	}
}

func TestHash(t *testing.T) {
	var (
		expect = "4528E6A7BB9341C36C425FAF40EF32C3"
	)
	hash, err := Md5("testdata/pass.md5")
	if err != nil {
		t.Error(err)
	}
	if expect != hash {
		t.Errorf("expected %s got %s", expect, hash)
	}
}

func TestExpand(t *testing.T) {
	expect := os.Getenv("HOME")
	if got := Path("$HOME").Expand(); got != expect {
		t.Errorf("expected %s got %s", expect, got)
	}
}

func TestJoin(t *testing.T) {
	expect := path.Join(os.Getenv("HOME"), "test")
	if got := Path("$HOME").Add("test"); got != expect {
		t.Errorf("expected %s got %s", expect, got)
	}
}

func testCat(t *testing.T) {
	err := Cat("file.go")
	if err != nil {
		t.Error(err)
	}
}
