package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type BenchmarkTest02233 struct{}

func (b *BenchmarkTest02233) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	r.ParseForm()
	param := r.FormValue("BenchmarkTest02233")
	bar := doSomething(param)

	w.Header().Set("X-XSS-Protection", "0")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(bar))
}

func doSomething(param string) string {
	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the last 'safe' value
	}
	return bar
}

func TestBenchmarkTest02233(t *testing.T) {
	request, err := http.NewRequest("POST", "/xss-04/BenchmarkTest02233", nil)
	if err != nil {
		t.Fatal(err)
	}

	query := request.URL.Query()
	query.Add("BenchmarkTest02233", "unsafeValue")
	request.URL.RawQuery = query.Encode()

	rr := httptest.NewRecorder()
	handler := &BenchmarkTest02233{}
	handler.ServeHTTP(rr, request)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "unsafeValue"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
