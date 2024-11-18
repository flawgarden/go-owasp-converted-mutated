package controllers

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

type BenchmarkTest01601Controller struct {
	http.Handler
}

func (c *BenchmarkTest01601Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		c.doPost(w, r)
	case http.MethodPost:
		c.doPost(w, r)
	default:
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
	}
}

func (c *BenchmarkTest01601Controller) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	values := r.URL.Query()["BenchmarkTest01601"]
	var param string
	if len(values) > 0 {
		param = values[0]
	}

	bar := new(Test).doSomething(r, param)

	var argList []string

	if os.Getenv("OS") == "Windows_NT" {
		argList = []string{"cmd.exe", "/c"}
	} else {
		argList = []string{"sh", "-c"}
	}
	argList = append(argList, fmt.Sprintf("echo %s", bar))

	cmd := exec.Command(argList[0], argList[1:]...)

	output, err := cmd.Output()
	if err != nil {
		http.Error(w, "Error executing command", http.StatusInternalServerError)
		return
	}
	w.Write(output)
}

type Test struct{}

func (t *Test) doSomething(r *http.Request, param string) string {
	bar := ""
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the param value
	}
	return bar
}
