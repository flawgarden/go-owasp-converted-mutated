package controllers

import (
	"fmt"
	"net/http"
	"os/exec"
)

type BenchmarkTest01779Controller struct {
	http.Handler
}

func (c *BenchmarkTest01779Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest01779")

	bar := new(Test).doSomething(r, param)

	var a1, a2 string
	if os := exec.Command("uname").Run(); os == nil {
		a1 = "sh"
		a2 = "-c"
	} else {
		a1 = "cmd.exe"
		a2 = "/c"
	}

	args := []string{a1, a2, fmt.Sprintf("echo %s", bar)}

	cmd := exec.Command(args[0], args[1:]...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		http.Error(w, "Error executing command", http.StatusInternalServerError)
		return
	}

	w.Write(output)
}

type Test struct{}

func (t *Test) doSomething(r *http.Request, param string) string {
	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[1]         // get the last 'safe' value
	}
	return bar
}
