package webserver

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndexHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()

	// Run handler
	handler := New()
	handler.ServeHTTP(rec, req)

	// Test
	if rec.Code != http.StatusOK {
		t.Errorf("Incorrect status code, got %v want %v", rec.Code, http.StatusOK)
	}
}
