package main

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/timenglesf/hivebox/internal/sensebox"
)

// newTestApplication creates a new instance of the application struct for
// testing purposes.
func newTestApplication(t *testing.T, setVersionEnvManually bool) (*application, error) {
	if !setVersionEnvManually {
		os.Setenv(VERSION_ENV, "1.0.0")
		defer os.Unsetenv(VERSION_ENV)
	}

	cfg, err := createConfig()
	if err != nil {
		return nil, err
	}

	return &application{
		logger:   slog.New(slog.NewTextHandler(io.Discard, nil)),
		cfg:      cfg,
		sensebox: &sensebox.SenseboxServiceMock{},
	}, nil
}

// testServer wraps an httptest.Server to provide additional methods for testing.
type testServer struct {
	*httptest.Server
}

// newTestServer creates a new testServer instance with the provided HTTP handler.
func newTestServer(t *testing.T, h http.Handler) *testServer {
	ts := httptest.NewServer(h)
	return &testServer{ts}
}

// get sends a GET request to the given URL path and unmarshals the JSON response into the provided interface.
// It returns the HTTP status code, response headers, and an error if any occurred during the request or unmarshaling.
func (ts *testServer) get(t *testing.T, urlPath string, v interface{}) (int, http.Header, error) {
	rs, err := ts.Client().Get(ts.URL + urlPath)
	if err != nil {
		t.Fatal(err)
	}
	defer rs.Body.Close()

	err = json.NewDecoder(rs.Body).Decode(v)
	if err != nil {
		return rs.StatusCode, rs.Header, err
	}

	return rs.StatusCode, rs.Header, nil
}
