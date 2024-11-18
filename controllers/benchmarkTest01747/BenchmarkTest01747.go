package controllers

import (
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

type BenchmarkTest01747 struct{}

func (b *BenchmarkTest01747) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest01747")
	bar := b.doSomething(r, param)
	startURIslashes := ""

	if strings.Contains(os.Getenv("OS"), "Windows") {
		startURIslashes = "/"
	} else {
		startURIslashes = "//"
	}

	fileURI, err := url.Parse(
		"file://" + startURIslashes + filepath.Join("TESTFILES_DIR", strings.ReplaceAll(bar, " ", "_")))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fileTarget := fileURI.Path
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	w.Write([]byte("Access to file: '" + encodeForHTML(fileTarget) + "' created.\n"))
	if _, err := os.Stat(fileTarget); os.IsNotExist(err) {
		w.Write([]byte("But file doesn't exist yet."))
	} else {
		w.Write([]byte("And file already exists."))
	}
}

func (b *BenchmarkTest01747) doSomething(r *http.Request, param string) string {
	bar := ""
	num := 86
	if (7*42)-num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}
	return bar
}

func encodeForHTML(str string) string {
	return strings.ReplaceAll(strings.ReplaceAll(str, "&", "&amp;"), "<", "&lt;")
}
