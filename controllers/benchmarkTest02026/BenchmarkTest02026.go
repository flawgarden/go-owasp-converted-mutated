package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
)

type BenchmarkTest02026 struct{}

func (bt *BenchmarkTest02026) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")

	param := ""
	headers := r.Header["BenchmarkTest02026"]
	if len(headers) > 0 {
		param = headers[0]
	}

	param, _ = url.QueryUnescape(param)

	bar := doSomething(r, param)

	fileTarget := fmt.Sprintf("%s/%s", os.Getenv("TESTFILES_DIR"), bar)
	w.Write([]byte(fmt.Sprintf("Access to file: '%s' created.", fileTarget)))
	if _, err := os.Stat(fileTarget); err == nil {
		w.Write([]byte(" And file already exists."))
	} else {
		w.Write([]byte(" But file doesn't exist yet."))
	}
}

func doSomething(r *http.Request, param string) string {
	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the first safe value
		bar = valuesList[0]         // get the last 'safe' value
	}
	return bar
}

func main() {
	http.Handle("/pathtraver-02/BenchmarkTest02026", &BenchmarkTest02026{})
	http.ListenAndServe(":8080", nil)
}
