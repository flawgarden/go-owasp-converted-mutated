//Original file region: null, null, null, null
//Mutated file region: null, null, null, null
//Semgrep original results: []
//Gosec original results: [78]
//CodeQL original results: [78]
//Snyk original results: [78]
//-------------
//Gosec analysis results: [78, 703]
//CodeQL analysis results: [78]
//Snyk analysis results: [1004]
//Semgrep analysis results: []
//Original file name: controllers/benchmarkTest01851/BenchmarkTest01851.go
//Original file CWE's: [78]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/varargs/varargs.tmt with name varargs_get_first_string_from_array_positive 
//Used extensions: 
//Program:
package controllers

import (
"net/http"
"os/exec"
"strings"
)

type BenchmarkTest01851 struct{}

func (b *BenchmarkTest01851) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	http.SetCookie(w, &http.Cookie{
		Name:   "BenchmarkTest01851",
		Value:  "ECHOOO",
		MaxAge: 60 * 3,
		Secure: true,
		Path:   r.URL.Path,
		Domain: r.Host,
	})
	http.ServeFile(w, r, "./cmdi-02/BenchmarkTest01851.html")
}

func (b *BenchmarkTest01851) Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	cookie, err := r.Cookie("BenchmarkTest01851")
	param := "noCookieValueSupplied"
	if err == nil {
		param = cookie.Value

param = getFirstStringFromArray(param, param)

	}

	bar := doSomething(param)

	var argList []string
	if strings.Contains(strings.ToLower(r.UserAgent()), "windows") {
		argList = []string{"cmd.exe", "/c", "echo " + bar}
	} else {
		argList = []string{"sh", "-c", "echo " + bar}
	}

	cmd := exec.Command(argList[0], argList[1:]...)
	output, err := cmd.Output()
	if err != nil {
		http.Error(w, "Error executing command", http.StatusInternalServerError)
		return
	}
	w.Write(output)
}

func doSomething(param string) string {
	bar := ""
	switchTarget := "C" // This simulates the switch case with a predetermined target
	switch switchTarget {
	case "A":
		bar = param
	case "B":
		bar = "bobs_your_uncle"
	case "C", "D":
		bar = param
	default:
		bar = "bobs_your_uncle"
	}
	return bar
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


