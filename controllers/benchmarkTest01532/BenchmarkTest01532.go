package controllers

import (
	"net/http"
	"os/exec"
)

type BenchmarkTest01532 struct{}

func (b *BenchmarkTest01532) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest01532")
	if param == "" {
		param = ""
	}

	bar := b.doSomething(param)

	var cmd string
	if os := exec.Command("uname"); os != nil {
		cmd = "echo " + bar
	}

	argsEnv := []string{"Foo=bar"}
	cmdExec := exec.Command("bash", "-c", cmd)
	cmdExec.Env = append(cmdExec.Env, argsEnv...)

	output, err := cmdExec.CombinedOutput()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(output)
}

func (b *BenchmarkTest01532) doSomething(param string) string {
	var bar string
	num := 86
	if (7*42)-num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}
	return bar
}

func main() {
	http.Handle("/cmdi-01/BenchmarkTest01532", &BenchmarkTest01532{})
	http.ListenAndServe(":8080", nil)
}
