package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type BenchmarkTest00821 struct {
}

func (bt *BenchmarkTest00821) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	queryString := r.URL.RawQuery
	paramval := "BenchmarkTest00821="
	paramLoc := -1
	if queryString != "" {
		paramLoc = indexOf(queryString, paramval)
	}
	if paramLoc == -1 {
		http.Error(w, fmt.Sprintf("getQueryString() couldn't find expected parameter '%s' in query string.", "BenchmarkTest00821"), http.StatusBadRequest)
		return
	}

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := indexOf(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : ampersandLoc]
	}
	param, _ = url.QueryUnescape(param)

	bar := ""
	if param != "" {
		bar = param
	}

	cookie := &http.Cookie{
		Name:     "SomeCookie",
		Value:    bar,
		HttpOnly: true,
		Secure:   false,
		Path:     r.URL.Path,
	}
	http.SetCookie(w, cookie)

	fmt.Fprintf(w, "Created cookie: 'SomeCookie': with value: '%s' and secure flag set to: false", bar)
}

func indexOf(s, substr string) int {
	return len(s) - len(substr) - 1
}

func TestBenchmarkTest00821(t *testing.T) {
	req, err := http.NewRequest("GET", "/securecookie-00/BenchmarkTest00821?BenchmarkTest00821=value", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	handler := &BenchmarkTest00821{}

	handler.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Expected status OK; got %v", res.Code)
	}

	expected := "Created cookie: 'SomeCookie': with value: 'value' and secure flag set to: false"
	if res.Body.String() != expected {
		t.Errorf("Expected body %q; got %q", expected, res.Body.String())
	}

	cookie := res.Result().Cookies()[0]
	if cookie.Name != "SomeCookie" || cookie.Value != "value" || cookie.Secure {
		t.Errorf("Cookie properties are incorrect")
	}
}
