package controllers

import (
	"fmt"
	"net/http"
	"os/exec"
	"strings"
)

type BenchmarkTest00573 struct{}

func (bt *BenchmarkTest00573) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		bt.doPost(w, r)
	} else if r.Method == http.MethodPost {
		bt.doPost(w, r)
	}
}

func (bt *BenchmarkTest00573) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	flag := true
	names := r.URL.Query()

	for name, values := range names {
		if flag {
			for _, value := range values {
				if value == "BenchmarkTest00573" {
					param = name
					flag = false
					break
				}
			}
		}
	}

	bar := param
	cmd := "your_command_here" // Replace with actual command
	argsEnv := []string{bar}
	command := exec.Command(cmd, argsEnv...)

	output, err := command.Output()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	escapedOutput := strings.ReplaceAll(string(output), "<", "&lt;")
	escapedOutput = strings.ReplaceAll(escapedOutput, ">", "&gt;")
	fmt.Fprintln(w, escapedOutput)
}
