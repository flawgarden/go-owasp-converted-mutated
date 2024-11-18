package controllers

import (
	"net/http"
	"net/url"
	"os"
)

type BenchmarkTest01746Controller struct {
}

func (c *BenchmarkTest01746Controller) Get(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest01746")
	bar := c.doSomething(r, param)

	startURIslashes := ""
	if os := os.Getenv("OS"); os != "" && os == "Windows_NT" {
		startURIslashes = "/"
	} else {
		startURIslashes = "//"
	}

	fileURI, err := url.Parse("file:" + startURIslashes + "/path/to/testfiles/" + bar)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fileTarget := fileURI.Path

	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	w.Write([]byte("Access to file: '" + fileTarget + "' created.\n"))
	if _, err := os.Stat(fileTarget); os.IsNotExist(err) {
		w.Write([]byte("But file doesn't exist yet.\n"))
	} else {
		w.Write([]byte("And file already exists.\n"))
	}
}

func (c *BenchmarkTest01746Controller) doSomething(r *http.Request, param string) string {
	bar := ""
	num := 106
	if (7*42)-num > 200 {
		bar = "This should never happen"
	} else {
		bar = param
	}
	return bar
}
