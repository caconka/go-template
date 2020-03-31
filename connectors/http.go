package connectors

import (
	"bytes"
	"fmt"
	"net/http"
)

func makeHTTPQuery(r *http.Request) (*http.Response, error) {

	client := &http.Client{}
	resp, err := client.Do(r)

	return resp, err
}

func newHTTPRequest(method, url string, body []byte, cookie *http.Cookie) (*http.Request, error) {
	r, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if nil != err {
		r = nil
	}

	if nil != cookie {
		r.AddCookie(cookie)
	}

	return r, err
}

func buildRetrieveEndpointURL(basePath, id string) string {
	return fmt.Sprintf("%s/j/%s", basePath, id)
}
