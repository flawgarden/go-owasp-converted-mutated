//Original file region: null, null, null, null
//Mutated file region: null, null, null, null
//Semgrep original results: []
//Gosec original results: [78]
//CodeQL original results: [78]
//Snyk original results: [78]
//-------------
//Gosec analysis results: [78, 703]
//CodeQL analysis results: [116, 79, 563]
//Snyk analysis results: []
//Semgrep analysis results: []
//Original file name: controllers/benchmarkTest02243/BenchmarkTest02243.go
//Original file CWE's: [78]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/field/constructors.tmt with name class_with_string_initialization_negative 
//Used extensions: 
//Program:
package controllers

import (
	"net/http"
	"os/exec"
)

type BenchmarkTest02243 struct{}

func (b *BenchmarkTest02243) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest02243")

	bar := doSomething(r, param)

	var argList []string

sh := NewStringHolder()
bar = sh.value

	if isWindows() {
		argList = []string{"cmd.exe", "/c", "echo " + bar}
	} else {
		argList = []string{"sh", "-c", "echo " + bar}
	}

	cmd := exec.Command(argList[0], argList[1:]...)
	output, err := cmd.Output()
	if err != nil {
		http.Error(w, "Problem executing command", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	w.Write(output)
}

func doSomething(r *http.Request, param string) string {
	bar := "safe!"
	map19941 := make(map[string]interface{})
	map19941["keyA-19941"] = "a-Value"
	map19941["keyB-19941"] = param
	map19941["keyC"] = "another-Value"
	bar = map19941["keyB-19941"].(string)

	return bar
}

func isWindows() bool {
	return exec.Command("cmd", "/c", "ver").Run() == nil
}