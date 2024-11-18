package controllers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"text/template"
)

func BenchmarkTest01924(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Referer", "http://example.com")
	doPost(w, r)
}

func doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	if referer := r.Header.Get("Referer"); referer != "" {
		param = referer
	}

	param, _ = url.QueryUnescape(param)

	bar := doSomething(r, param)

	w.Header().Set("X-XSS-Protection", "0")
	_, _ = w.Write([]byte(bar))
}

func doSomething(r *http.Request, param string) string {
	return htmlEscape(param)
}

func htmlEscape(param string) string {
	return template.HTMLEscapeString(param)
}

func TestBenchmark(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost/xss-03/BenchmarkTest01924", nil)
	rr := httptest.NewRecorder()

	BenchmarkTest01924(rr, req)

	res := rr.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK; got %v", res.Status)
	}
}
