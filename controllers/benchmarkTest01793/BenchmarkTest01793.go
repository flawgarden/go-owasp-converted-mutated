package controllers

import (
	"fmt"
	"net/http"
	"os/exec"
)

type BenchmarkTest01793 struct{}

func (b *BenchmarkTest01793) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		b.doPost(w, r)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func (b *BenchmarkTest01793) doPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest01793")

	bar := b.doSomething(param)

	cmd := getInsecureOSCommandString()
	args := []string{cmd}
	argsEnv := []string{bar}

	process := exec.Command(args[0], args[1:]...)
	process.Env = append(process.Env, argsEnv...)

	output, err := process.Output()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(output)
}

func (b *BenchmarkTest01793) doSomething(param string) string {
	var bar string
	guess := "ABC"
	switchTarget := guess[1] // condition 'B', which is safe

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

	return bar
}

func getInsecureOSCommandString() string {
	return "your_command_here"
}

func main() {
	http.Handle("/cmdi-02/BenchmarkTest01793", &BenchmarkTest01793{})
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
