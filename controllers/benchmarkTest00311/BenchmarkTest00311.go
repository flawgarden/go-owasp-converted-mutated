package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"os/exec"
)

type BenchmarkTest00311 struct{}

func (b *BenchmarkTest00311) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	if headers := r.Header["BenchmarkTest00311"]; len(headers) > 0 {
		param = headers[0]
	}

	param, _ = url.QueryUnescape(param)

	bar := ""
	guess := "ABC"
	switchTarget := guess[2]

	switch switchTarget {
	case 'A':
		bar = param
	case 'B':
		bar = "bobs_your_uncle"
	case 'C', 'D':
		bar = param
	default:
		bar = "bobs_your_uncle"
	}

	cmd := getInsecureOSCommandString()
	argsEnv := []string{bar}

	cmdExec := exec.Command(cmd, argsEnv...)
	output, err := cmdExec.CombinedOutput()
	if err != nil {
		fmt.Println("Problem executing cmdi - TestCase")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(output)
}

func getInsecureOSCommandString() string {
	return "echo"
}

func main() {
	http.Handle("/cmdi-00/BenchmarkTest00311", &BenchmarkTest00311{})
	http.ListenAndServe(":8080", nil)
}
