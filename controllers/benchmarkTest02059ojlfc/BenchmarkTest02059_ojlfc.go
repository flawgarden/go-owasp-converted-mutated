//Semgrep original results: []
//Gosec original results: [78]
//CodeQL original results: [78]
//Snyk original results: [78]
//-------------
//Gosec analysis results: [676, 78, 703]
//CodeQL analysis results: [116, 79]
//Snyk analysis results: []
//Semgrep analysis results: []
//Original file name: controllers/benchmarkTest02059/BenchmarkTest02059.go
//Original file CWE's: [78]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/namedreturns.tmt with name named_return_swap_simple_negative 
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

a, _ := swap(bar, "suffix")
bar = a

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


