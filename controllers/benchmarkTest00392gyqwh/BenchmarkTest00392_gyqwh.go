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
//Original file name: controllers/benchmarkTest00392/BenchmarkTest00392.go
//Original file CWE's: [79]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/virtuality/interface.tmt with name impl_binary_op_interface_class2_negative 
//Used extensions: 
//Program:
package controllers

import (
	"net/http"
	"strings"
)

type BenchmarkTest00392 struct{}

func (b *BenchmarkTest00392) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := r.URL.Query().Get("BenchmarkTest00392")
	if param == "" {
		param = ""
	}

	bar := param
	if param != "" && len(param) > 1 {
		sbxyz38384 := strings.Builder{}
		sbxyz38384.WriteString(param)
		bar = sbxyz38384.String()[:len(param)-1] + "Z"
	}

	w.Header().Set("X-XSS-Protection", "0")
	length := 1
	if bar != "" {

var a12341 BinaryOpInterface = &ImplBinaryOpInterfaceClass2{}
bar = a12341.InterfaceCall(bar, "")

		length = len(bar)
		w.Write([]byte(bar[:length]))
	}
}

func main() {
	http.Handle("/xss-00/BenchmarkTest00392", &BenchmarkTest00392{})
	http.ListenAndServe(":8080", nil)
}
