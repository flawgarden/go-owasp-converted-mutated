package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

type BenchmarkTest struct{}

func (b *BenchmarkTest) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	param := r.Header.Get("BenchmarkTest01029")
	param, _ = url.QueryUnescape(param)

	bar := b.doSomething(r, param)

	startURIslashes := ""
	if os.PathSeparator == '\\' {
		startURIslashes = "\\"
	} else {
		startURIslashes = "//"
	}

	fileURI := fmt.Sprintf("file://%s%s", startURIslashes, filepath.Join("testfiles", bar))
	fileTarget := filepath.FromSlash(fileURI)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Access to file: '%s' created.\n", fileTarget)
	if _, err := os.Stat(fileTarget); os.IsNotExist(err) {
		fmt.Fprintln(w, "But file doesn't exist yet.")
	} else {
		fmt.Fprintln(w, "And file already exists.")
	}
}

func (b *BenchmarkTest) doSomething(r *http.Request, param string) string {
	return param // Simple pass-through for this example, actual logic can be added
}
