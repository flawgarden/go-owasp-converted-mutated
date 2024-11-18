package controllers

import (
	"fmt"
	"net/http"
	"os/exec"
	"strings"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

func init() {
	// Database initialization code here
}

type BenchmarkTest00658 struct {
}

func (b *BenchmarkTest00658) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest00658")
	if param == "" {
		param = ""
	}

	var bar string
	num := 106

	if (7*18)+num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}

	cmd := ""
	if strings.Contains(strings.ToLower(r.UserAgent()), "windows") {
		cmd = "echo "
	}

	command := cmd + bar

	if err := executeCommand(command, w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func executeCommand(command string, w http.ResponseWriter) error {
	cmd := exec.Command("bash", "-c", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error executing command: %v", err)
	}
	w.Write(output)
	return nil
}
