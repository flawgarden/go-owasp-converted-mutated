//Original file region: null, null, null, null
//Mutated file region: null, null, null, null
//Semgrep original results: []
//Gosec original results: [78]
//CodeQL original results: [78]
//Snyk original results: [78]
//-------------
//Gosec analysis results: [676, 78, 703]
//CodeQL analysis results: [116, 79]
//Snyk analysis results: [79, 78]
//Semgrep analysis results: []
//Original file name: controllers/benchmarkTest01285/BenchmarkTest01285.go
//Original file CWE's: [78]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/virtuality/interface.tmt with name impl_binary_op_interface_class2_negative 
//Used extensions: 
//Program:
package controllers

import (
	"net/http"
	"os/exec"
)

type BenchmarkTest01285 struct {
}

func (b *BenchmarkTest01285) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest01285")
	if param == "" {
		param = ""
	}

	bar := b.doSomething(param)

	var cmd string

var a12341 BinaryOpInterface = &ImplBinaryOpInterfaceClass2{}
bar = a12341.InterfaceCall(bar, "")

	osName := "UNIX"
	if osName == "Windows" {
		cmd = `echo `
	}

	command := exec.Command(cmd + bar)
	output, err := command.Output()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	w.Write(output)
}

func (b *BenchmarkTest01285) doSomething(param string) string {
	var bar string
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the param value
	}
	return bar
}

func main() {
	http.Handle("/cmdi-01/BenchmarkTest01285", &BenchmarkTest01285{})
	http.ListenAndServe(":8080", nil)
}
