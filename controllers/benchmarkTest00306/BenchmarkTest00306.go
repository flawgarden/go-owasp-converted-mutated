package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"os/exec"
)

type BenchmarkTest00306 struct{}

func (bt *BenchmarkTest00306) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	if headers := r.Header["BenchmarkTest00306"]; len(headers) > 0 {
		param = headers[0]
	}

	param, _ = url.QueryUnescape(param)

	var bar string
	num := 196
	if (500/42)+num > 200 {
		bar = param
	} else {
		bar = "This should never happen"
	}

	cmd := "your_command_here" // <-- Replace with actual command
	args := []string{cmd}
	argsEnv := []string{bar}

	command := exec.Command(args[0], args[1:]...)
	command.Env = append(os.Environ(), argsEnv...)

	output, err := command.CombinedOutput()
	if err != nil {
		fmt.Println("Problem executing cmdi - TestCase")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(output)
}

func main() {
	http.Handle("/cmdi-00/BenchmarkTest00306", &BenchmarkTest00306{})
	http.ListenAndServe(":8080", nil)
}
