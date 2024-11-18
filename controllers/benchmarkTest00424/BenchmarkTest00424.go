package controllers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"
)

type BenchmarkTest00424 struct{}

func (b *BenchmarkTest00424) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest00424")
	if param == "" {
		param = ""
	}

	bar := ""
	if param != "" {
		decodedBytes, err := base64.StdEncoding.DecodeString(base64.StdEncoding.EncodeToString([]byte(param)))
		if err == nil {
			bar = string(decodedBytes)
		}
	}

	r.AddCookie(&http.Cookie{Name: "userid", Value: bar})

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Item: 'userid' with value: '%s' saved in session.", htmlEscape(bar))
}

func htmlEscape(s string) string {
	return template.HTMLEscapeString(s)
}

func TestBenchmark00424(t *testing.T) {
	req := httptest.NewRequest("GET", "/trustbound-00/BenchmarkTest00424?BenchmarkTest00424=test", nil)
	w := httptest.NewRecorder()

	handler := &BenchmarkTest00424{}
	handler.ServeHTTP(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %v", res.StatusCode)
	}

	var responseBody string
	json.NewDecoder(res.Body).Decode(&responseBody)
	expectedResponse := "Item: 'userid' with value: 'test' saved in session."
	if responseBody != expectedResponse {
		t.Errorf("Expected response body '%s', got '%s'", expectedResponse, responseBody)
	}
}
