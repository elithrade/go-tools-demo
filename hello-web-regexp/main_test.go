// Package main provides ...
package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	cases := []struct {
		in, out string
	}{
		{"helloworld@gmail.com", "gmail user helloworld"},
		{"helloworld@live.com", "helloworld"},
		{"helloworld", "helloworld"},
	}

	for _, c := range cases {
		req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/"+c.in, nil)
		if err != nil {
			t.Fatalf("Could not create request: %v", err)
		}
		// ResponseRecorder is an implementation of http.ResponseWriter that
		// records its mutations for later inspection in tests.
		w := httptest.NewRecorder()
		handler(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status %d, but git %d", http.StatusOK, w.Code)
		}
		if !strings.Contains(w.Body.String(), c.out) {
			t.Errorf("Unexpected body response %q", w.Body.String())
		}
	}
}
