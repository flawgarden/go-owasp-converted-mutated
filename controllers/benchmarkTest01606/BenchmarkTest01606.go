package controllers

import (
	"fmt"
	"net/http"
	"os/exec"
	"runtime"
)

type BenchmarkTest01606 struct{}

func (b *BenchmarkTest01606) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	values := r.URL.Query()["BenchmarkTest01606"]
	var param string
	if len(values) > 0 {
		param = values[0]
	} else {
		param = ""
	}

	bar := b.doSomething(param)

	var cmd string
	if isWindows() {
		cmd = "cmd /C echo " // for Windows
	} else {
		cmd = "/bin/echo " // for Unix-based
	}

	output, err := executeCommand(cmd + bar)
	if err != nil {
		fmt.Println("Problem executing cmdi - TestCase")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(output))
}

func (b *BenchmarkTest01606) doSomething(param string) string {
	var bar string
	num := 86
	if (7*42)-num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}
	return bar
}

func isWindows() bool {
	return runtime.GOOS == "windows"
}

func executeCommand(cmd string) (string, error) {
	out, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
