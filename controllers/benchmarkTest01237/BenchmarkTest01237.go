package controllers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

type BenchmarkTest01237 struct{}

func (b *BenchmarkTest01237) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		b.doPost(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (b *BenchmarkTest01237) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.URL.Query().Get("BenchmarkTest01237")
	if param == "" {
		param = ""
	}

	bar := b.doSomething(param)

	startURIslashes := ""
	if os.PathSeparator == '\\' {
		startURIslashes = "//"
	} else {
		startURIslashes = "/"
	}

	fileURI := fmt.Sprintf("file:%s%s", startURIslashes, filepath.Join(os.Getenv("TESTFILES_DIR"), bar))
	fileTarget := filepath.Clean(fileURI)

	fmt.Fprintf(w, "Access to file: '%s' created.", fileTarget)
	if _, err := os.Stat(fileTarget); err == nil {
		fmt.Fprintln(w, " And file already exists.")
	} else {
		fmt.Fprintln(w, " But file doesn't exist yet.")
	}
}

func (b *BenchmarkTest01237) doSomething(param string) string {
	bar := ""
	num := 86
	if (7*42)-num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}
	return bar
}

func main() {
	http.Handle("/pathtraver-01/BenchmarkTest01237", &BenchmarkTest01237{})
	http.ListenAndServe(":8080", nil)
}
