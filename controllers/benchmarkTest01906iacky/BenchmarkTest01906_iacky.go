//Original file region: null, null, null, null
//Mutated file region: null, null, null, null
//Semgrep original results: []
//Gosec original results: [22]
//CodeQL original results: [22]
//Snyk original results: [22]
//-------------
//Gosec analysis results: [22, 276]
//CodeQL analysis results: [563]
//Snyk analysis results: []
//Semgrep analysis results: []
//Original file name: controllers/benchmarkTest01906/BenchmarkTest01906.go
//Original file CWE's: [22]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/field/nested.tmt with name nested_field_depth_3_negative 
//Used extensions: 
//Program:
package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

type BenchmarkTest01906Controller struct {
	http.Handler
}

func (c *BenchmarkTest01906Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		c.doPost(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (c *BenchmarkTest01906Controller) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")

	param := r.Header.Get("BenchmarkTest01906")
	param, _ = url.QueryUnescape(param)

nested7231 := NewNestedFields3("oWCGQ")
param = nested7231.nested1.nested1.nested1.value

	bar := doSomething(param)
	fileName := fmt.Sprintf("/path/to/directory/%s", bar)

	fos, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("Couldn't open FileOutputStream on file: '%s'\n", fileName)
		return
	}
	defer fos.Close()

	_, err = w.Write([]byte(fmt.Sprintf("Now ready to write to file: %s", htmlEscape(fileName))))
	if err != nil {
		fmt.Printf("Error writing response: %v\n", err)
	}
}

func doSomething(param string) string {
	bar := ""
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the param value
	}
	return bar
}

func htmlEscape(s string) string {
	return jsonEscape(s)
}

func jsonEscape(s string) string {
	output, _ := json.Marshal(s)
	return string(output)
}
