//Original file region: null, null, null, null
//Mutated file region: null, null, null, null
//Semgrep original results: []
//Gosec original results: [78]
//CodeQL original results: [78]
//Snyk original results: []
//-------------
//Gosec analysis results: [78, 703]
//CodeQL analysis results: [116, 79, 563]
//Semgrep analysis results: []
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest00304/BenchmarkTest00304.go
//Original file CWE's: [78]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/collections/list.tmt with name list_set_negative 
//Used extensions: MACRO_Create_List -> ~[MACRO_ListName]~ := make([] ~[TYPE@1]~, 0) | MACRO_Add_CONST_ToList -> ~[MACRO_ListName]~ = append(~[MACRO_ListName]~, ~[CONST_~[TYPE@1]~@1]~) | MACRO_ListName -> list787231 | MACRO_ListName -> list787231 | MACRO_ListName -> list787231 | MACRO_ListName -> list787231
//Program:
package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"
)

type BenchmarkTest00304 struct{}

func (b *BenchmarkTest00304) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.Header.Get("BenchmarkTest00304")
	param, _ = url.QueryUnescape(param)

list787231 := make([] string, 0)
list787231 = append(list787231, "tvIvM")
param = list787231[0]

	var bar string
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the param value
	}

	var cmd string
	if strings.Contains(os.Getenv("OS"), "Windows") {
		cmd = fmt.Sprintf("cmd.exe /c echo %s", bar)
	} else {
		cmd = fmt.Sprintf("sh -c ls %s", bar)
	}

	output, err := exec.Command("bash", "-c", cmd).CombinedOutput()
	if err != nil {
		fmt.Fprintf(w, "Problem executing cmdi - TestCase: %s", err.Error())
		return
	}
	w.Write(output)
}