//Original file region: null, null, null, null
//Mutated file region: null, null, null, null
//Semgrep original results: [79]
//Gosec original results: []
//CodeQL original results: [79]
//Snyk original results: [79]
//-------------
//Semgrep analysis results: [79, 319]
//CodeQL analysis results: []
//Snyk analysis results: [79]
//Gosec analysis results: []
//Original file name: controllers/benchmarkTest01267/BenchmarkTest01267.go
//Original file CWE's: [79]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/field/nested.tmt with name nested_field_depth_3_array_positive 
//Used extensions: 
//Program:
package controllers

import (
	"fmt"
	"net/http"
)

type BenchmarkTest01267 struct{}

func (b *BenchmarkTest01267) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	r.ParseForm()
	param := r.FormValue("BenchmarkTest01267")
	if param == "" {
		param = ""
	}

	bar := new(Test).doSomething(r, param)

arr4124 := []string{bar}
nested7231 := NewNestedFields3FromArray(arr4124)
bar = nested7231.nested1.nested1.nested1.values[0]

	w.Header().Set("X-XSS-Protection", "0")
	w.Write([]byte(bar))
}

type Test struct{}

func (t *Test) doSomething(r *http.Request, param string) string {
	bar := "safe!"
	map53210 := make(map[string]interface{})
	map53210["keyA-53210"] = "a-Value"
	map53210["keyB-53210"] = param
	map53210["keyC"] = "another-Value"
	bar = map53210["keyB-53210"].(string)

	return bar
}

func main() {
	http.Handle("/xss-02/BenchmarkTest01267", &BenchmarkTest01267{})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
