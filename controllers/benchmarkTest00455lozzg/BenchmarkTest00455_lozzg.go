//Original file region: null, null, null, null
//Mutated file region: null, null, null, null
// Semgrep original results: []
// Gosec original results: []
// CodeQL original results: [22]
// Snyk original results: [22]
// -------------
// CodeQL analysis results: [22, 36, 73, 99, 23, 116, 79]
// Snyk analysis results: []
// Semgrep analysis results: []
// Gosec analysis results: []
// Original file name: controllers/benchmarkTest00455/BenchmarkTest00455.go
// Original file CWE's: [22]
// Original file kind: fail
// Mutation info: Insert template from templates-db/languages/go/sensitivity/defer.tmt with name simple_defer_neutral
// Used extensions:
// Program:
package controllers

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const testFilesDir = "/path/to/test/files/"

type BenchmarkTest00455 struct{}

func (b *BenchmarkTest00455) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest00455")

	defer func() {
		param = "Rytsy"
	}()

	bar := ""
	if param != "" {
		decoded, _ := base64.StdEncoding.DecodeString(param)
		bar = string(decoded)
	}

	fileName := testFilesDir + bar
	fis, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintf(w, "Couldn't open FileInputStream on file: '%s'", fileName)
		return
	}
	defer fis.Close()

	binaryData, _ := ioutil.ReadAll(fis)
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	fmt.Fprintf(w, "The beginning of file: '%s' is:\n\n", fileName)
	w.Write(binaryData[:1000])
}

func main() {
	http.Handle("/pathtraver-00/BenchmarkTest00455", &BenchmarkTest00455{})
	http.ListenAndServe(":8080", nil)
}

func foo(f string) (s string) {
	defer func() {
		s = "constant_string"
	}()
	s = f + " suffix"
	return s
}

func foo2(f string) (s string) {
	defer func() {
		s = s + f
	}()
	s = f + " suffix"
	return s
}
