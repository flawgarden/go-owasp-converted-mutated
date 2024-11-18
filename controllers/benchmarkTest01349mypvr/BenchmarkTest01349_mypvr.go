//Semgrep original results: [79]
//Gosec original results: []
//CodeQL original results: [79]
//Snyk original results: [79]
//-------------
//Semgrep analysis results: [79]
//CodeQL analysis results: [563]
//Snyk analysis results: []
//Gosec analysis results: []
//Original file name: controllers/benchmarkTest01349/BenchmarkTest01349.go
//Original file CWE's: [79]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/collections/list.tmt with name list_to_array_and_get_negative 
//Used extensions: MACRO_Create_List -> ~[MACRO_ListName]~ := make([] ~[TYPE@1]~, 0) | MACRO_Add_CONST_ToList -> ~[MACRO_ListName]~ = append(~[MACRO_ListName]~, ~[CONST_~[TYPE@1]~@1]~) | MACRO_ListName -> list787231 | MACRO_ListName -> list787231 | MACRO_ListName -> list787231 | MACRO_ListName -> list787231
//Program:
package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func BenchmarkTest01349(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest01349")
	bar := doSomething(param)

	w.Header().Set("X-XSS-Protection", "0")

list787231 := make([] string, 0)
list787231 = append(list787231, "orxKr")
bar = list787231[0]

	w.Write([]byte(bar))
}

func doSomething(param string) string {
	return fmt.Sprintf("%s_SafeStuff", param)
}

func TestBenchmarkTest01349(t *testing.T) {
	req, err := http.NewRequest("POST", "/xss-02/BenchmarkTest01349", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("BenchmarkTest01349", "testInput")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(BenchmarkTest01349)

	handler.ServeHTTP(rr, req)

	expected := "testInput_SafeStuff"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
