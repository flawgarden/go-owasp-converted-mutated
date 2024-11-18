package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"os/exec"
	"strings"
)

type BenchmarkTest00814 struct{}

func (bt *BenchmarkTest00814) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	queryString := r.URL.RawQuery
	paramval := "BenchmarkTest00814="
	paramLoc := strings.Index(queryString, paramval)
	if paramLoc == -1 {
		http.Error(w, fmt.Sprintf("getQueryString() couldn't find expected parameter '%s' in query string.", "BenchmarkTest00814"), http.StatusBadRequest)
		return
	}

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := strings.Index(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : paramLoc+ampersandLoc]
	}
	param, _ = url.QueryUnescape(param)

	bar := param
	num := 106

	if (7*18)+num > 200 {
		bar = "This_should_always_happen"
	}

	var argList []string

	if strings.Contains(strings.ToLower(r.UserAgent()), "windows") {
		argList = []string{"cmd.exe", "/c", "echo " + bar}
	} else {
		argList = []string{"sh", "-c", "echo " + bar}
	}

	cmd := exec.Command(argList[0], argList[1:]...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		http.Error(w, "Problem executing cmdi - ProcessBuilder Test Case", http.StatusInternalServerError)
		return
	}

	w.Write(output)
}

func main() {
	http.Handle("/cmdi-00/BenchmarkTest00814", &BenchmarkTest00814{})
	http.ListenAndServe(":8080", nil)
}
