package betfair

import (
	"bytes"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"net/http"
	"testing"
)

var credentials = InteractiveCredentials{
	Username: "joe",
	Password: "password",
	Key:      "key-secret",
}

type roundTripFunc func(req *http.Request) *http.Response

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func newTestClient(fn roundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
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

type MockStore struct {
	mock.Mock
}

func (m *MockStore) Set(ctx context.Context, key, value string) error {
	args := m.Called(ctx, key, value)
	return args.Error(0)
}

func (m *MockStore) Get(ctx context.Context, key string) (string, error) {
	args := m.Called(ctx, key)
	return args.Get(0).(string), args.Error(1)
}
