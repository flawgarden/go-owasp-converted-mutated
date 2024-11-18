package controllers

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

type BenchmarkTest00659 struct{}

func (b *BenchmarkTest00659) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest00659")
	if param == "" {
		param = ""
	}

	var bar string
	guess := "ABC"
	switchTarget := guess[1]

	switch switchTarget {
	case 'A':
		bar = param
	case 'B':
		bar = "bob"
	case 'C', 'D':
		bar = param
	default:
		bar = "bob's your uncle"
	}

	cmd := ""
	if os.Getenv("OS") == "Windows_NT" {
		cmd = "cmd /c echo "
	}

	// Execution of the command
	fullCmd := cmd + bar
	command := exec.Command("cmd", "/c", fullCmd)
	output, err := command.CombinedOutput()

	if err != nil {
		fmt.Fprintln(w, "Problem executing cmdi - TestCase")
		fmt.Fprintln(w, err.Error())
		return
	}

	w.Write(output)
}
