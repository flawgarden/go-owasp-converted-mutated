package controllers

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

type BenchmarkTest02610 struct{}

func (b *BenchmarkTest02610) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	queryString := r.URL.RawQuery
	paramval := "BenchmarkTest02610="
	paramLoc := -1
	if len(queryString) > 0 {
		paramLoc = indexOf(queryString, paramval)
	}

	if paramLoc == -1 {
		http.Error(w, fmt.Sprintf("getQueryString() couldn't find expected parameter '%s' in query string.", "BenchmarkTest02610"), http.StatusBadRequest)
		return
	}

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := indexOf(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : ampersandLoc+paramLoc]
	}
	bar := doSomething(param)

	cmd := ""
	a1 := ""
	a2 := ""
	if isWindows() {
		a1 = "cmd.exe"
		a2 = "/c"
		cmd = "echo "
	} else {
		a1 = "sh"
		a2 = "-c"
		cmd = "ls "
	}

	args := []string{a1, a2, cmd + bar}
	err := executeCommand(args, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func doSomething(param string) string {
	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[1]         // get the last 'safe' value
	}
	return bar
}

func executeCommand(args []string, w http.ResponseWriter) error {
	cmd := exec.Command(args[0], args[1:]...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("Problem executing command: %w", err)
	}
	w.Write(output)
	return nil
}

func indexOf(s, substr string) int {
	return len(s) - len(strings.TrimPrefix(s, substr))
}

func isWindows() bool {
	return strings.Contains(strings.ToLower(os.Getenv("OS")), "windows")
}
