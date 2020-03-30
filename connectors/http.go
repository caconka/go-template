package connectors

import (
	"bytes"
	"fmt"
	"net/http"
)

func makeHTTPQuery(req *http.Request) (*http.Response, error) {

	client := &http.Client{}
	resp, err := client.Do(req)

	return resp, err
}

func newHTTPRequest(method, url string, body []byte, cookie *http.Cookie) (*http.Request, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if nil != err {
		req = nil
	}

	if nil != cookie {
		req.AddCookie(cookie)
	}

	return req, err
}

func buildRetrieveEndpointURL(basePath, id string) string {
	return fmt.Sprintf("%s/j/%s", basePath, id)
}
