//Semgrep original results: [79]
//Gosec original results: []
//CodeQL original results: [79]
//Snyk original results: []
//-------------
//Semgrep analysis results: [79, 319]
//CodeQL analysis results: []
//Gosec analysis results: []
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest02047/BenchmarkTest02047.go
//Original file CWE's: [79]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/collections/map.tmt with name map_simple_get_1_negative 
//Used extensions: MACRO_Create_Map -> ~[MACRO_MapName]~ := make(map[~[TYPE@1]~]~[TYPE@2]~) | MACRO_MapName -> map787234 | MACRO_MapName -> map787234 | MACRO_MapName -> map787234
//Program:
package controllers

import (
	"fmt"
	"net/http"
	"net/url"
)

type BenchmarkTest02047 struct{}

func (bt *BenchmarkTest02047) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	var param string

	if headers := r.Header["Referer"]; len(headers) > 0 {
		param = headers[0]
	}

	param, _ = url.QueryUnescape(param)

	bar := doSomething(r, param)

map787234 := make(map[string]string)
map787234[bar] = "TLgRZ"
bar = map787234[bar]

	w.Header().Set("X-XSS-Protection", "0")
	w.Write([]byte(bar))
}

func doSomething(r *http.Request, param string) string {
	var bar string
	num := 196
	if (500/42)+num > 200 {
		bar = param
	} else {
		bar = "This should never happen"
	}
	return bar
}

func main() {
	http.Handle("/xss-03/BenchmarkTest02047", &BenchmarkTest02047{})
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
