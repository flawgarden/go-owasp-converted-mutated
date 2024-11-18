package controllers

import (
	"net/http"
	"os/exec"
)

type BenchmarkTest01600Controller struct {
	http.Handler
}

func (c *BenchmarkTest01600Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	values := r.Form["BenchmarkTest01600"]
	var param string
	if len(values) > 0 {
		param = values[0]
	} else {
		param = ""
	}

	bar := new(Test).doSomething(r, param)

	var argList []string
	if isWindows() {
		argList = append(argList, "cmd.exe", "/c")
	} else {
		argList = append(argList, "sh", "-c")
	}
	argList = append(argList, "echo "+bar)

	cmd := exec.Command(argList[0], argList[1:]...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		http.Error(w, "Problem executing command", http.StatusInternalServerError)
		return
	}
	w.Write(out)
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

func isWindows() bool {
	return false // implement your OS detection logic here
}
