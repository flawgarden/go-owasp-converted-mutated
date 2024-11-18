package controllers

import (
	"net/http"
	"os/exec"
)

type BenchmarkTest02516 struct{}

func (b *BenchmarkTest02516) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.URL.Query().Get("BenchmarkTest02516")
	bar := doSomething(param)

	cmd := "some-command" // Replace with actual command
	argsEnv := []string{bar}

	if err := execCommand(cmd, argsEnv, w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func doSomething(param string) string {
	bar := "safe!"
	map74796 := make(map[string]interface{})
	map74796["keyA-74796"] = "a-Value"
	map74796["keyB-74796"] = param
	map74796["keyC"] = "another-Value"
	bar = map74796["keyB-74796"].(string)

	return bar
}

func execCommand(cmd string, args []string, w http.ResponseWriter) error {
	r := exec.Command(cmd, args...)
	output, err := r.CombinedOutput()
	if err != nil {
		return err
	}
	_, _ = w.Write(output)
	return nil
}
