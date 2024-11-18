//Semgrep original results: [79]
//Gosec original results: []
//CodeQL original results: [79]
//Snyk original results: [79]
//-------------
//Semgrep analysis results: [79, 319]
//CodeQL analysis results: [116, 79]
//Snyk analysis results: []
//Gosec analysis results: []
//Original file name: controllers/benchmarkTest00392/BenchmarkTest00392.go
//Original file CWE's: [79]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/varargs/varargs.tmt with name varargs_get_string_with_index_positive 
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

bar = getStringWithIndex(0, bar, param)

		length = len(bar)
		w.Write([]byte(bar[:length]))
	}
}

func main() {
	http.Handle("/xss-00/BenchmarkTest00392", &BenchmarkTest00392{})
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


