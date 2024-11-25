//Original file region: null, null, null, null
//Mutated file region: null, null, null, null
//Semgrep original results: [79]
//Gosec original results: []
//CodeQL original results: [79]
//Snyk original results: [79]
//-------------
//Semgrep analysis results: [79]
//CodeQL analysis results: [563]
//Snyk analysis results: []
//Gosec analysis results: []
//Original file name: controllers/benchmarkTest02132/BenchmarkTest02132.go
//Original file CWE's: [79]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/namedreturns.tmt with name named_return_zero_values_negative 
//Used extensions: 
//Program:
package controllers

import (
	"fmt"
	"net/http"

	"xorm.io/xorm"
)

type BenchmarkTest02132 struct {
	engine *xorm.Engine
}

func NewBenchmarkTest02132(engine *xorm.Engine) *BenchmarkTest02132 {
	return &BenchmarkTest02132{engine: engine}
}

func (bt *BenchmarkTest02132) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		bt.doPost(w, r)
	} else if r.Method == http.MethodPost {
		bt.doPost(w, r)
	}
}

func (bt *BenchmarkTest02132) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.URL.Query().Get("BenchmarkTest02132")

a123, _ := getZeroValues()
param = a123

	if param == "" {
		param = ""
	}

	bar := doSomething(param)

	w.Header().Set("X-XSS-Protection", "0")
	w.Write([]byte(bar))
}

func doSomething(param string) string {
	a76744 := param
	b76744 := fmt.Sprintf("%s SafeStuff", a76744)
	b76744 = b76744[:len(b76744)-5] + "Chars"
	map76744 := map[string]string{"key76744": b76744}
	c76744 := map76744["key76744"]
	d76744 := c76744[:len(c76744)-1]
	e76744 := string([]byte(d76744)) // Base64 decoding removed for simplicity

	f76744 := e76744 // Base64 encoding/decoding handled here
	thing := CreateThing()
	bar := thing.DoSomething(f76744)

	return bar
}

// Mocking the ThingInterface for illustration purposes
type ThingInterface interface {
	DoSomething(input string) string
}

func CreateThing() ThingInterface {
	return &mockThing{}
}

type mockThing struct{}

func (m *mockThing) DoSomething(input string) string {
	return input
}

func concat(a string, b string) (res string) {
    res = a + b
    return
}

func swap(a string, b string) (first string, second string) {
	first, second = b, a
	return
}

func brokenConcat(a string, b string) (result string) {
	defer func() {
		result = b
	}()
	result = a + b
	return
}

func getZeroValues() (x string, y string) {
    return
}


