package controllers

import (
	"net/http"
	"net/url"
	"os/exec"
)

type BenchmarkTest01939 struct {
	http.Handler
}

func (b *BenchmarkTest01939) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.Header.Get("BenchmarkTest01939")
	param, _ = url.QueryUnescape(param)


	cmd := "insecure_command" // Replace with actual command retrieval logic
	args := []string{cmd}

	out, err := exec.Command(args[0], args...).CombinedOutput()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(out)
}

func doSomething(r *http.Request, param string) string {
	var bar string
	num := 86
	if (7*42)-num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}
	return bar
}
