package goutils

import (
	"io/ioutil"
	"os"
)

// WriteFile write bytes to file with perm of 0644
func WriteFile(filename string, data []byte) bool {
	if data != nil {
		err := ioutil.WriteFile(filename, data, 0644)
		if err != nil {
			return false
		}
		return true
	}
	return false
}

// ReadFile read bytes from file
func ReadFile(filename string) (string, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(bytes), err
}

// IsFileExists check whether file is exists
func IsFileExists(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
