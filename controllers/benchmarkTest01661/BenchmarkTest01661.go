package controllers

import (
	"fmt"
	"html"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

type BenchmarkTest01661 struct{}

func (b *BenchmarkTest01661) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	queryString := r.URL.RawQuery
	paramVal := "BenchmarkTest01661="
	paramLoc := -1
	if queryString != "" {
		paramLoc = strings.Index(queryString, paramVal)
	}
	if paramLoc == -1 {
		w.Write([]byte(fmt.Sprintf("getQueryString() couldn't find expected parameter '%s' in query string.", "BenchmarkTest01661")))
		return
	}

	param := queryString[paramLoc+len(paramVal):]
	ampersandLoc := strings.Index(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramVal) : paramLoc+ampersandLoc]
	}
	param, _ = url.QueryUnescape(param)

	bar := new(Test).doSomething(param)

	w.Header().Set("X-XSS-Protection", "0")
	w.Write([]byte(fmt.Sprintf("Formatted like: %s and %s.", "a", bar)))
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	bar := html.EscapeString(param)
	return bar
}

func TestBenchmark(t *testing.T) {
	req, err := http.NewRequest("GET", "/xss-03/BenchmarkTest01661?BenchmarkTest01661=test", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc((&BenchmarkTest01661{}).ServeHTTP)

	handler.ServeHTTP(rr, req)

	expected := "Formatted like: a and test."
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
