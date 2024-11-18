package controllers

import (
	"net/http"
	"net/url"
	"os/exec"
)

type BenchmarkTest01190 struct{}

func (b *BenchmarkTest01190) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	headers := r.Header["BenchmarkTest01190"]
	if len(headers) > 0 {
		param = headers[0]
	}

	param, _ = url.QueryUnescape(param)

	bar := b.doSomething(r, param)

	cmd := "your-command" // Укажите команду, которую нужно выполнить
	args := []string{cmd}
	argsEnv := []string{bar}

	if err := executeCommand(args, argsEnv, w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (b *BenchmarkTest01190) doSomething(r *http.Request, param string) string {
	num := 106
	if (7*42)-num > 200 {
		return "This should never happen"
	}
	return param
}

func executeCommand(args, argsEnv []string, w http.ResponseWriter) error {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Env = argsEnv
	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	_, _ = w.Write(output)
	return nil
}
