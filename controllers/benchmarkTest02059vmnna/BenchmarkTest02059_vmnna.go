//Original file region: null, null, null, null
//Mutated file region: null, null, null, null
//Semgrep original results: []
//Gosec original results: [78]
//CodeQL original results: [78]
//Snyk original results: [78]
//-------------
//Gosec analysis results: [676, 78, 703]
//CodeQL analysis results: [78, 116, 79]
//Snyk analysis results: []
//Semgrep analysis results: []
//Original file name: controllers/benchmarkTest02059/BenchmarkTest02059.go
//Original file CWE's: [78]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/defer.tmt with name simple_non_defer_negative 
//Used extensions: 
//Program:
package controllers

import (
	"net/http"
	"os/exec"
	"strings"
)

type BenchmarkTest02059 struct{}

func (b *BenchmarkTest02059) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.Header.Get("BenchmarkTest02059")

func() {
    param = "FJeEs"
}()

	param = decode(param)

	bar := doSomething(r, param)

	var a1, a2 string
	if strings.Contains(strings.ToLower(r.UserAgent()), "windows") {
		a1 = "cmd.exe"
		a2 = "/c"
	} else {
		a1 = "sh"
		a2 = "-c"
	}
	args := []string{a1, a2, "echo " + bar}

	cmd := exec.Command(args[0], args[1:]...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		http.Error(w, "Error executing command", http.StatusInternalServerError)
		return
	}
	w.Write(output)
}

func doSomething(r *http.Request, param string) string {
	bar := "safe!"
	collection := map[string]interface{}{
		"keyA-36421": "a-Value",
		"keyB-36421": param,
		"keyC":       "another-Value",
	}
	bar = collection["keyB-36421"].(string)
	return bar
}

func decode(value string) string {
	if len(value) == 0 {
		return ""
	}
	decoded, err := exec.Command("echo", value).Output()
	if err != nil {
		return value
	}
	return string(decoded)
}

func main() {
	http.Handle("/cmdi-02/BenchmarkTest02059", &BenchmarkTest02059{})
	http.ListenAndServe(":8080", nil)
}

func foo(f string) (s string) {
	defer func() {
		s = "constant_string"
	}()
	s = f + " suffix"
	return s
}

func foo2(f string) (s string) {
	defer func() {
		s = s + f
	}()
	s = f + " suffix"
	return s
}


