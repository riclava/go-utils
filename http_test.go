package goutils

import "testing"

// TestGetHTTP test GetHTTP()
func TestGetHTTP(t *testing.T) {
	headers := make(map[string]string)
	body, err := HTTPGet("https://github.com", headers)
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log(body)
	}
}
