package goutils

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

// HTTPPost send data to server with http POST method and return server payload, if not ok return error
func HTTPPost(url string, headers map[string]string, data string) (string, error) {
	return doHTTP(url, headers, "POST")
}

// HTTPPut send data to server with http PUT method and return server payload, if not ok return error
func HTTPPut(url string, headers map[string]string, data string) (string, error) {
	return doHTTP(url, headers, "PUT")
}

// HTTPGet get data from server with http PUT method and return server payload, if not ok return error
func HTTPGet(url string, headers map[string]string) (string, error) {
	return doHTTP(url, headers, "GET")
}

// HTTPDelete delete resource of server with http DELETE method, if not ok return error
func HTTPDelete(url string, headers map[string]string) (string, error) {
	return doHTTP(url, headers, "DELETE")
}

func doHTTP(url string, headers map[string]string, method string) (string, error) {
	client := http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return "", err
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if !(resp.StatusCode >= 200 && resp.StatusCode < 300) {
		return "", errors.New("Http response code not valid [" + strconv.Itoa(resp.StatusCode) + "]")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
