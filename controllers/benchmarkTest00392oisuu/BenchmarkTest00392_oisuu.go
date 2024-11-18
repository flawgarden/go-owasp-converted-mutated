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
//Mutation info: Insert template from templates-db/languages/go/sensitivity/closure.tmt with name closure_with_outer_variable_positive 
//Used extensions: 
//Program:
package controllers

import (
"net/http"
"strings"
"fmt"
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

stringReturner := func() func() string {
    return func() string {
        return bar
    }
}

stringRet := stringReturner()
bar = stringRet()

	w.Header().Set("X-XSS-Protection", "0")
	length := 1
	if bar != "" {
		length = len(bar)
		w.Write([]byte(bar[:length]))
	}
}

func main() {
	http.Handle("/xss-00/BenchmarkTest00392", &BenchmarkTest00392{})
	http.ListenAndServe(":8080", nil)
}

func makePrefixer(prefix string) func(string) string {
    return func(value string) string {
        return fmt.Sprintf("%s%s", prefix, value)
    }
}

func makeMessageGenerator(prefix string) func() func(string) string {
	return func() func(string) string {
		return func(name string) string {
			return prefix + name
		}
	}
}

func makeMessageGeneratorBroken(prefix string) func() func(string) string {
	return func() func(string) string {
		return func(name string) string {
			return name
		}
	}
}


