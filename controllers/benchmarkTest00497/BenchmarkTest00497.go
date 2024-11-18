package controllers

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

type BenchmarkTest00497 struct{}

func (b *BenchmarkTest00497) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest00497")
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

	cmd := "/path/to/command" // пример команды
	args := []string{cmd}
	argsEnv := []string{bar}

	command := exec.Command(args[0], args[1:]...)
	command.Env = append(os.Environ(), argsEnv...)

	output, err := command.CombinedOutput()
	if err != nil {
		http.Error(w, fmt.Sprintf("Problem executing cmd: %v", err), http.StatusInternalServerError)
		return
	}
	w.Write(output)
}
