//Original file region: null, null, null, null
//Mutated file region: null, null, null, null
//Semgrep original results: [79]
//Gosec original results: []
//CodeQL original results: []
//Snyk original results: [79]
//-------------
//Semgrep analysis results: [79]
//Snyk analysis results: []
//Gosec analysis results: []
//CodeQL analysis results: []
//Original file name: controllers/benchmarkTest00473/BenchmarkTest00473.go
//Original file CWE's: [79]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/field/staticField.tmt with name class_with_static_string_field_negative 
//Used extensions: 
//Program:
package controllers

import (
	"encoding/base64"
	"fmt"
	"net/http"
)

type BenchmarkTest00473 struct{}

func (b *BenchmarkTest00473) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "0")
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.URL.Query().Get("BenchmarkTest00473")
	bar := ""
	if param != "" {
		decoded, _ := base64.StdEncoding.DecodeString(base64.StdEncoding.EncodeToString([]byte(param)))
		bar = string(decoded)

sfh := NewStaticFieldHolder()
bar = sfh.value

	}

	obj := []interface{}{"a", "b"}
	w.Write([]byte(fmt.Sprintf(bar, obj...)))
}
