package controllers

import (
	"fmt"
	"net/http"
	"os"
)

type BenchmarkTest01329 struct{}

func (bt *BenchmarkTest01329) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest01329")
	bar := bt.doSomething(param)

	startURIslashes := ""
	if os.PathSeparator == '\\' {
		startURIslashes = "//"
	} else {
		startURIslashes = "/"
	}

	fileURI := fmt.Sprintf("file:%s%s", startURIslashes, bar)
	fileTarget := fileURI

	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	fmt.Fprintf(w, "Access to file: '%s' created.", fileTarget)
	if _, err := os.Stat(fileTarget); os.IsNotExist(err) {
		fmt.Fprintln(w, " But file doesn't exist yet.")
	} else {
		fmt.Fprintln(w, " And file already exists.")
	}
}

func (bt *BenchmarkTest01329) doSomething(param string) string {
	bar := ""
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = append(valuesList[:0], valuesList[1:]...) // remove the 1st safe value
		bar = valuesList[0]                                    // get the param value
	}
	return bar
}
