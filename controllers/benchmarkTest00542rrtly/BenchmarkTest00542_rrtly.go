//Original file region: null, null, null, null
//Mutated file region: null, null, null, null
//Semgrep original results: [79]
//Gosec original results: []
//CodeQL original results: []
//Snyk original results: [79]
//-------------
//Semgrep analysis results: [319]
//Snyk analysis results: [79]
//Gosec analysis results: []
//CodeQL analysis results: []
//Original file name: controllers/benchmarkTest00542/BenchmarkTest00542.go
//Original file CWE's: [79]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/concurrency/concurrency.tmt with name thread_set_positive 
//Used extensions: 
//Program:
package controllers

import (
"fmt"
"net/http"
"sync"
)

type BenchmarkTest00542 struct{}

func (b *BenchmarkTest00542) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		b.doPost(w, r)
		return
	}
	b.doPost(w, r)
}

func (b *BenchmarkTest00542) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	flag := true
	names := r.URL.Query()

	for name, values := range names {
		if flag {
			for _, value := range values {
				if value == "BenchmarkTest00542" {
					param = name
					flag = false
					break
				}
			}
		}
	}

	var bar string
	guess := "ABC"
	switchTarget := guess[2]

	switch switchTarget {
	case 'A':
		bar = param
	case 'B':

w := &Wrapper[string]{Value: guess}
task := NewSettingTask(w, guess)
var wg sync.WaitGroup
wg.Add(1)
go func() {
    defer wg.Done()
    task.Run()
}()
wg.Wait()
guess = w.Value

		bar = "bobs_your_uncle"
	case 'C', 'D':
		bar = param
	default:
		bar = "bobs_your_uncle"
	}

	w.Header().Set("X-XSS-Protection", "0")
	if _, err := fmt.Fprintf(w, bar); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.Handle("/xss-00/BenchmarkTest00542", &BenchmarkTest00542{})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
