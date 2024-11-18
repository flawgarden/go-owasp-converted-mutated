//Semgrep original results: []
//Gosec original results: [22]
//CodeQL original results: [22]
//Snyk original results: []
//-------------
//Gosec analysis results: [676, 22, 703]
//CodeQL analysis results: []
//Semgrep analysis results: []
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest00362/BenchmarkTest00362.go
//Original file CWE's: [22]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/closure.tmt with name closure_with_outer_variable_positive 
//Used extensions: 
//Program:
package controllers

import (
"fmt"
"net/http"
"os"
"path/filepath"
)

type BenchmarkTest00362 struct{}

func (b *BenchmarkTest00362) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	param := r.URL.Query().Get("BenchmarkTest00362")
	if param == "" {
		param = ""
	}

	bar := param
	fileName := filepath.Join("testfiles", bar)

stringReturner := func() func() string {
    return func() string {
        return fileName
    }
}

stringRet := stringReturner()
fileName = stringRet()

	var fos *os.File

	defer func() {
		if fos != nil {
			fos.Close()
		}
	}()

	fos, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Couldn't open FileOutputStream on file: '" + fileName + "'")
		return
	}

	_, _ = w.Write([]byte("Now ready to write to file: " + fileName))
}

func main() {
	http.Handle("/pathtraver-00/BenchmarkTest00362", &BenchmarkTest00362{})
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


