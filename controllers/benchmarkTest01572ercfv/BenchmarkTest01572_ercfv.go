//Original file region: null, null, null, null
//Mutated file region: null, null, null, null
//Semgrep original results: []
//Gosec original results: [22]
//CodeQL original results: [22]
//Snyk original results: [22]
//-------------
//Gosec analysis results: [22]
//CodeQL analysis results: [22, 36, 73, 99, 23, 116, 79]
//Snyk analysis results: []
//Semgrep analysis results: []
//Original file name: controllers/benchmarkTest01572/BenchmarkTest01572.go
//Original file CWE's: [22]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/varargs/varargs.tmt with name varargs_with_generics_positive 
//Used extensions: 
//Program:
package controllers

import (
"fmt"
"net/http"
"os"
"path/filepath"
"strings"
)

type BenchmarkTest01572Controller struct {
	http.Handler
}

func (c *BenchmarkTest01572Controller) Get(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()["BenchmarkTest01572"]
	var param string
	if len(values) > 0 {
		param = values[0]
	}

	bar := new(Test).doSomething(r, param)

bar = varargsWithGenerics(bar, "Goodh")

	fileName := filepath.Join("testfiles", bar)
	fis, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintf(w, "Problem getting FileInputStream: %s", err.Error())
		return
	}
	defer fis.Close()

	b := make([]byte, 1000)
	size, _ := fis.Read(b)
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	fmt.Fprintf(w, "The beginning of file: '%s' is:\n\n%s", htmlEscape(fileName), htmlEscape(string(b[:size])))
}

func htmlEscape(str string) string {
	return strings.ReplaceAll(strings.ReplaceAll(str, "&", "&amp;"), "<", "&lt;")
}

type Test struct{}

func (t *Test) doSomething(r *http.Request, param string) string {
	thing := newThing()
	bar := thing.doSomething(param)
	return bar
}

func newThing() ThingInterface {
	return &Thing{}
}

type ThingInterface interface {
	doSomething(string) string
}

type Thing struct{}

func (t *Thing) doSomething(param string) string {
	return param // Здесь может быть другая логика обработки
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


