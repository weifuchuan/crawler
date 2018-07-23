package util

import (
	"io/ioutil"
	"net/http"
)

// Get for use http get method.
func Get(url string) (res []byte, err error) {
	r, err := http.Get(url)
	if err != nil {
		return
	}
	defer r.Body.Close()
	return ioutil.ReadAll(r.Body)
}
