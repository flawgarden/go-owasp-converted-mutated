//Semgrep original results: [79]
//Gosec original results: []
//CodeQL original results: [79]
//Snyk original results: [79]
//-------------
//Semgrep analysis results: [79]
//CodeQL analysis results: [563]
//Snyk analysis results: []
//Gosec analysis results: []
//Original file name: controllers/benchmarkTest01173/BenchmarkTest01173.go
//Original file CWE's: [79]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/closure.tmt with name closure_with_outer_variable_negative 
//Used extensions: 
//Program:
package controllers

import (
"net/http"
"net/url"
_ "github.com/go-sql-driver/mysql"
"fmt"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01173 struct{}

func (bt *BenchmarkTest01173) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.Header.Get("X-XSS-Protection")
	param := r.Header.Get("Referer")

	param, _ = url.QueryUnescape(param)

	bar := new(Test).doSomething(r, param)

	w.Header().Set("X-XSS-Protection", "0")

stringReturner := func() func() string {
    return func() string {
        return "hBnXU"
    }
}

stringRet := stringReturner()
bar = stringRet()

	w.Write([]byte(bar))
}

type Test struct{}

func (t *Test) doSomething(r *http.Request, param string) string {
	bar := ""
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value

		bar = valuesList[0] // get the param value
	}
	return bar
}

func init() {
	http.Handle("/xss-02/BenchmarkTest01173", &BenchmarkTest01173{})
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


