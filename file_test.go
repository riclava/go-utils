package goutils

import (
	"testing"
)

// TestReadWriteFile test WriteFile() and ReadFile()
func TestReadWriteFile(t *testing.T) {
	file := "./debug.log"
	data := "THIS IS TEST DATA"
	if WriteFile(file, []byte(data)) {
		t.Log("write file success")
	} else {
		t.Error("write file failed")
	}
	if IsFileExists(file) {
		newData, err := ReadFile("./debug.log")
		if err != nil {
			t.Error(err)
		} else {
			if newData != data {
				t.Error("data isn't match")
			} else {
				t.Log("read file success")
			}
		}
	} else {
		t.Error("can not find the file we've already written")
	}
}
