//Semgrep original results: [79]
//Gosec original results: []
//CodeQL original results: [79]
//Snyk original results: []
//-------------
//Semgrep analysis results: [79]
//CodeQL analysis results: []
//Gosec analysis results: []
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest00643/BenchmarkTest00643.go
//Original file CWE's: [79]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/types/assertions.tmt with name type_assertion_in_switch_negative 
//Used extensions: 
//Program:
package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type BenchmarkTest00643 struct{}

func (h *BenchmarkTest00643) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest00643")
	if param == "" {
		param = ""
	}

	sbxyz31207 := strings.Builder{}
	sbxyz31207.WriteString(param)
	bar := sbxyz31207.String() + "_SafeStuff"

var i interface{} = 1562581557
switch i.(type) {
case int:
    bar = "ofeFB"
case string:
    bar = bar + "Dojcg"
default:
    bar = "culCM"
}

	w.Header().Set("X-XSS-Protection", "0")
	w.Write([]byte(bar))
}

func TestBenchmarkTest00643(t *testing.T) {
	handler := &BenchmarkTest00643{}
	req := httptest.NewRequest(http.MethodPost, "/xss-01/BenchmarkTest00643", nil)
	req.Form.Add("BenchmarkTest00643", "test")
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	expected := "test_SafeStuff"
	if rr.Body.String() != expected {
		t.Errorf("expected %v, got %v", expected, rr.Body.String())
	}
}
