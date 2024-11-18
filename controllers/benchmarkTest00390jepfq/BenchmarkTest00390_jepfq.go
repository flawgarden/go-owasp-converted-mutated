//Semgrep original results: [79]
//Gosec original results: []
//CodeQL original results: []
//Snyk original results: [79]
//-------------
//Semgrep analysis results: [79, 319]
//Snyk analysis results: []
//Gosec analysis results: []
//CodeQL analysis results: []
//Original file name: controllers/benchmarkTest00390/BenchmarkTest00390.go
//Original file CWE's: [79]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/field/constructors.tmt with name class_with_instance_initializer_positive 
//Used extensions: 
//Program:
package controllers

import (
	"encoding/base64"
	"net/http"
)

type BenchmarkTest00390 struct{}

func (b *BenchmarkTest00390) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	w.Header().Set("X-XSS-Protection", "0")

	param := r.URL.Query().Get("BenchmarkTest00390")
	if param == "" {
		param = ""
	}

ii := NewInstanceInitializer(param)
param = ii.list[1]

	var bar string
	if param != "" {
		decoded, _ := base64.StdEncoding.DecodeString(base64.StdEncoding.EncodeToString([]byte(param)))
		bar = string(decoded)
	}

	w.Write([]byte(bar))
}

func main() {
	http.Handle("/xss-00/BenchmarkTest00390", &BenchmarkTest00390{})
	http.ListenAndServe(":8080", nil)
}
