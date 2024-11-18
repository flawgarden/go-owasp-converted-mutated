package controllers

import (
	"net/http"
	"os/exec"
)

type BenchmarkTest02342 struct{}

func (b *BenchmarkTest02342) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := ""
	flag := true
	for name, values := range r.Form {
		if flag {
			for _, value := range values {
				if value == "BenchmarkTest02342" {
					param = name
					flag = false
					break
				}
			}
		}
	}

	bar := doSomething(r, param)

	cmd := ""
	if cmd = "echo"; isWindows() {
		cmd = "cmd.exe"
	}

	cmdOut, err := exec.Command(cmd, bar).CombinedOutput()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(cmdOut)
}

func doSomething(r *http.Request, param string) string {
	return param
}

func isWindows() bool {
	return false
}
