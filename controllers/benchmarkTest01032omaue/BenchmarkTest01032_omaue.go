//Original file region: null, null, null, null
//Mutated file region: null, null, null, null
//Semgrep original results: []
//Gosec original results: [22]
//CodeQL original results: [22]
//Snyk original results: []
//-------------
//Gosec analysis results: [676, 22, 703]
//CodeQL analysis results: [563]
//Semgrep analysis results: []
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest01032/BenchmarkTest01032.go
//Original file CWE's: [22]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/field/nested.tmt with name nested_field_depth_4_negative 
//Used extensions: 
//Program:
package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
)

type BenchmarkTest01032 struct{}

func (b *BenchmarkTest01032) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.Header.Get("BenchmarkTest01032")
	param, _ = url.QueryUnescape(param)

	bar := b.doSomething(r, param)

	var fileName string

nested7231 := NewNestedFields4("pahuQ")
bar = nested7231.nested1.nested1.nested1.nested1.value

	var fos *os.File
	defer func() {
		if fos != nil {
			fos.Close()
		}
	}()

	fileName = "/path/to/test/files/" + bar

	fos, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Couldn't open FileOutputStream on file: '" + fileName + "'")
		return
	}

	w.Write([]byte("Now ready to write to file: " + fileName))
}

func (b *BenchmarkTest01032) doSomething(r *http.Request, param string) string {
	bar := ""

	num := 196
	if (500/42)+num > 200 {
		bar = param
	} else {
		bar = "This should never happen"
	}

	return bar
}

func main() {
	http.Handle("/pathtraver-01/BenchmarkTest01032", &BenchmarkTest01032{})
	http.ListenAndServe(":8080", nil)
}
