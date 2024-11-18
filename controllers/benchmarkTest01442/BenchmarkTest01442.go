package controllers

import (
	"net/http"
	"os/exec"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01442 struct {
}

func (b *BenchmarkTest01442) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		b.doPost(w, r)
		return
	}
	if r.Method == http.MethodPost {
		b.doPost(w, r)
		return
	}
}

func (b *BenchmarkTest01442) doPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := ""
	flag := true
	for name, values := range r.Form {
		if flag {
			for _, value := range values {
				if value == "BenchmarkTest01442" {
					param = name
					flag = false
					break
				}
			}
		}
	}

	bar := b.doSomething(r, param)

	cmd := getInsecureOSCommandString()
	args := []string{cmd}
	argsEnv := []string{bar}

	cmdExec := exec.Command(args[0], args[1:]...)
	cmdExec.Env = append(cmdExec.Env, argsEnv...)

	output, err := cmdExec.CombinedOutput()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(output)
}

func (b *BenchmarkTest01442) doSomething(r *http.Request, param string) string {
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

	return bar
}

func getInsecureOSCommandString() string {
	return "insecure-command" // example command
}
