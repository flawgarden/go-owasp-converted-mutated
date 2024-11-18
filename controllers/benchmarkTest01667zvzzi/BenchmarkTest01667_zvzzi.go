//Semgrep original results: [79]
//Gosec original results: []
//CodeQL original results: [79]
//Snyk original results: [79]
//-------------
//Semgrep analysis results: [79, 319]
//CodeQL analysis results: []
//Snyk analysis results: []
//Gosec analysis results: []
//Original file name: controllers/benchmarkTest01667/BenchmarkTest01667.go
//Original file CWE's: [79]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/types/embedding.tmt with name embedding_anonymous_fields_negative 
//Used extensions: 
//Program:
package controllers

import (
	"net/http"
)

type BenchmarkTest01667 struct{}

func (b *BenchmarkTest01667) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	queryString := r.URL.Query()
	param := queryString.Get("BenchmarkTest01667")
	if param == "" {
		http.Error(w, "getQueryString() couldn't find expected parameter 'BenchmarkTest01667' in query string.", http.StatusBadRequest)
		return
	}

	bar := b.doSomething(param)

a := AnonFieldHolder{
    Value2: "const_str",
    Anon: Anon{
        Value1: bar,
    },
}
bar = a.Value2

	w.Header().Set("X-XSS-Protection", "0")
	w.Write([]byte(bar))
}

func (b *BenchmarkTest01667) doSomething(param string) string {
	bar := "safe!"
	map96771 := make(map[string]interface{})
	map96771["keyA-96771"] = "a-Value"
	map96771["keyB-96771"] = param
	map96771["keyC"] = "another-Value"
	bar, _ = map96771["keyB-96771"].(string)
	return bar
}

func main() {
	http.Handle("/xss-03/BenchmarkTest01667", &BenchmarkTest01667{})
	http.ListenAndServe(":8080", nil)
}
