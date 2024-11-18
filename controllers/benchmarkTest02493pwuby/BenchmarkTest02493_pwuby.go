//Semgrep original results: [79]
//Gosec original results: []
//CodeQL original results: [79]
//Snyk original results: [79]
//-------------
//Semgrep analysis results: [79, 319]
//CodeQL analysis results: []
//Snyk analysis results: [79]
//Gosec analysis results: []
//Original file name: controllers/benchmarkTest02493/BenchmarkTest02493.go
//Original file CWE's: [79]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/concurrency/channels.tmt with name channel_struct_data_int_negative 
//Used extensions: 
//Program:
package controllers

import (
"net/http"
"sync"
)

type BenchmarkTest02493 struct{}

func (b *BenchmarkTest02493) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	if r.Method == http.MethodGet {
		b.doGet(w, r)
	} else if r.Method == http.MethodPost {
		b.doPost(w, r)
	}
}

func (b *BenchmarkTest02493) doGet(w http.ResponseWriter, r *http.Request) {
	b.doPost(w, r)
}

func (b *BenchmarkTest02493) doPost(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()["BenchmarkTest02493"]
	var param string
	if len(values) > 0 {
		param = values[0]
	} else {
		param = ""
	}

	bar := doSomething(param)

dataChannel := make(chan DataInt, 1)
dataChannel <- DataInt{Value: 1}

var wg sync.WaitGroup
wg.Add(1)

go func() {
    defer wg.Done()
    data := <-dataChannel
    data.Value += 1
    dataChannel <- data
}()

wg.Wait()

result := <-dataChannel
if (result.Value == 2) {
    bar = "constant_string"
}

	w.Header().Set("X-XSS-Protection", "0")
	if bar != "" {
		w.Write([]byte(bar))
	}
}

func doSomething(param string) string {
	bar := "safe!"
	map4720 := make(map[string]interface{})
	map4720["keyA-4720"] = "a-Value"
	map4720["keyB-4720"] = param
	map4720["keyC"] = "another-Value"
	bar = map4720["keyB-4720"].(string)

	return bar
}

func main() {
	http.Handle("/xss-05/BenchmarkTest02493", &BenchmarkTest02493{})
	http.ListenAndServe(":8080", nil)
}
