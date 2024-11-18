package controllers

import (
	"fmt"
	"net/http"
	"os/exec"
	"runtime"
)

type BenchmarkTest00572 struct{}

func (b *BenchmarkTest00572) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	response := "text/html;charset=UTF-8"
	w.Header().Set("Content-Type", response)

	var param string
	flag := true
	for name, values := range r.Form {
		if flag {
			for _, value := range values {
				if value == "BenchmarkTest00572" {
					param = name
					flag = false
				}
			}
		}
	}

	bar := ""
	num := 106
	if (7*18)+num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}

	cmd := ""
	if os := runtime.GOOS; os == "windows" {
		cmd = "cmd.exe /C echo " // platform-specific command
	} else {
		cmd = "echo "
	}

	command := exec.Command(cmd + bar)
	output, err := command.CombinedOutput()
	if err != nil {
		http.Error(w, fmt.Sprintf("Problem executing cmd - TestCase: %v", err), http.StatusInternalServerError)
		return
	}
	w.Write(output)
}
