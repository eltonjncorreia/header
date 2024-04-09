package header

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHeader(t *testing.T) {
	exampleController := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	request := httptest.NewRequest("GET", "http://example.com/", nil)

	write := httptest.NewRecorder()

	Headers(exampleController).ServeHTTP(write, request)

	expectedHeaders := map[string]string{
		"Content-Security-Policy": "default-src 'self'; style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com",
		"Referrer-Policy":         "origin-when-cross-origin",
		"X-Content-Type-Options":  "nosniff",
		"X-Frame-Options":         "deny",
		"X-XSS-Protection":        "0",
		"Server":                  "Go",
	}

	for key, value := range expectedHeaders {
		if write.Header().Get(key) != value {
			t.Errorf("Expected header %s to be %s, got %s", key, value, write.Header().Get(key))
		}
	}
}
