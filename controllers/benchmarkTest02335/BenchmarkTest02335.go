package controllers

import (
	"fmt"
	"net/http"
	"os/exec"
	"runtime"
)

type BenchmarkTest02335 struct{}

func (bt *BenchmarkTest02335) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	flag := true
	for name, values := range r.URL.Query() {
		if flag {
			for _, value := range values {
				if value == "BenchmarkTest02335" {
					param = name
					flag = false
					break
				}
			}
		}
	}

	bar := doSomething(r, param)

	var a1, a2 string
	if isWindows() {
		a1 = "cmd.exe"
		a2 = "/c"
	} else {
		a1 = "sh"
		a2 = "-c"
	}
	args := []string{a1, a2, fmt.Sprintf("echo %s", bar)}

	cmd := exec.Command(args[0], args[1:]...)
	output, err := cmd.Output()
	if err != nil {
		http.Error(w, "Problem executing command", http.StatusInternalServerError)
		return
	}
	w.Write(output)
}

func doSomething(r *http.Request, param string) string {
	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[1]         // get the last 'safe' value
	}
	return bar
}

func isWindows() bool {
	return runtime.GOOS == "windows"
}
