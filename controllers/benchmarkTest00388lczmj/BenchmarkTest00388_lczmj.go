//Original file region: null, null, null, null
//Mutated file region: null, null, null, null
//Semgrep original results: [79]
//Gosec original results: []
//CodeQL original results: [79]
//Snyk original results: []
//-------------
//Semgrep analysis results: [79, 319]
//CodeQL analysis results: [563]
//Gosec analysis results: []
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest00388/BenchmarkTest00388.go
//Original file CWE's: [79]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/varargs/varargs.tmt with name varargs_get_first_string_negative 
//Used extensions: 
//Program:
package controllers

import (
"fmt"
"net/http"
"strings"
)

func BenchmarkTest00388(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := r.URL.Query().Get("BenchmarkTest00388")

param = getFirstString("QZNOm", "kAABh")

	if param == "" {
		param = ""
	}

	sbxyz30382 := strings.Builder{}
	sbxyz30382.WriteString(param)
	bar := sbxyz30382.String() + "_SafeStuff"

	w.Header().Set("X-XSS-Protection", "0")
	fmt.Fprintln(w, []rune(bar))
}

func main() {
	http.HandleFunc("/xss-00/BenchmarkTest00388", BenchmarkTest00388)
	http.ListenAndServe(":8080", nil)
}

func getFirstString(lines ...string) string {
    return getStringWithIndex(0, lines...)
}

func getStringWithIndex(ind int, lines ...string) string {
    return lines[ind]
}

func getFirstStringFromArray(lines ...string) string {
    return lines[0]
}

func varargsWithGenerics[T any](elements ...T) T {
    return elements[0]
}

func combineStrings(strs ...string) string {
    return strings.Join(strs, ", ")
}


