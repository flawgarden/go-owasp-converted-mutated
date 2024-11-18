package controllers

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

type BenchmarkTest01944 struct{}

func (bt *BenchmarkTest01944) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.Header.Get("BenchmarkTest01944")
	bar := doSomething(param)

	var cmd string
	if isWindows() {
		cmd = "echo " + bar
	}

	argsEnv := []string{"Foo=bar"}
	cmdExec := exec.Command("cmd", "/C", cmd)

	for _, env := range argsEnv {
		cmdExec.Env = append(cmdExec.Env, env)
	}

	output, err := cmdExec.Output()
	if err != nil {
		fmt.Println("Problem executing cmdi - TestCase")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(output)
}

func doSomething(param string) string {
	var bar string
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

func isWindows() bool {
	return os.Getenv("OS") == "Windows_NT"
}
