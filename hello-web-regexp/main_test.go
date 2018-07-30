// Package main provides ...
package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/helloworld@gmail.com", nil)
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
	if !strings.Contains(w.Body.String(), "gmail user") {
		t.Errorf("Unexpected body response %q", w.Body.String())
	}
}
