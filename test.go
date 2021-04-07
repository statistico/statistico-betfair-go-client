package betfair

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)


var credentials = InteractiveCredentials{
	Username: "joe",
	Password: "password",
	Key:      "key-secret",
	Token:    "token",
}

// roundTripFunc .
type roundTripFunc func(req *http.Request) *http.Response

// roundTrip .
func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func newTestClient(fn roundTripFunc) *http.Client {
	return &http.Client{
		Transport: roundTripFunc(fn),
	}
}

func mockResponseServer(t *testing.T, body string, code int, url string) *http.Client {
	return newTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, url, req.URL.String())

		return &http.Response{
			StatusCode: code,
			Body:       ioutil.NopCloser(bytes.NewBufferString(body)),
		}
	})
}
