//Original file region: null, null, null, null
//Mutated file region: null, null, null, null
//Semgrep original results: [79]
//Gosec original results: []
//CodeQL original results: [79]
//Snyk original results: [79]
//-------------
//Semgrep analysis results: [79]
//CodeQL analysis results: []
//Snyk analysis results: [79]
//Gosec analysis results: []
//Original file name: controllers/benchmarkTest01349/BenchmarkTest01349.go
//Original file CWE's: [79]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/collections/map.tmt with name map_remove_1_negative 
//Used extensions: MACRO_Create_Map -> ~[MACRO_MapName]~ := make(map[~[TYPE@1]~]~[TYPE@2]~) | MACRO_Add_EXPR_ToMap -> ~[MACRO_MapName]~[~[EXPR_~[TYPE@1]~@1]~] = ~[EXPR_~[TYPE@2]~@2]~ | MACRO_MapName -> map787234 | MACRO_MapName -> map787234 | MACRO_MapName -> map787234 | MACRO_MapName -> map787234 | MACRO_MapName -> map787234 | MACRO_MapName -> map787234
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

map787234 := make(map[int]string)
map787234[1344077265] = "LoFvl"
map787234[1344077265] = bar
if _, ok := map787234[1344077265]; ok {
    delete(map787234, 1344077265)
}
value7843, exists := map787234[1344077265]
if !exists {
    value7843 = "RMJkV"
}
bar = value7843

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
