package utils

import (
	"bytes"
	"net/http"
)

func MakeHTTPQuery(r *http.Request) (*http.Response, error) {

	client := &http.Client{}
	resp, err := client.Do(r)

	return resp, err
}

func NewHTTPRequest(method, url string, body []byte, cookie *http.Cookie) (*http.Request, error) {
	r, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if nil != err {
		r = nil
	}

	if nil != cookie {
		r.AddCookie(cookie)
	}

	return r, err
}
