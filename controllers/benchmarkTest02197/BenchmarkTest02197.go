package controllers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
)

type BenchmarkTest02197 struct{}

func (b *BenchmarkTest02197) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest02197")
	bar := doSomething(param)

	fileTarget := filepath.Join(os.Getenv("TESTFILES_DIR"), bar)
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	fmt.Fprintf(w, "Access to file: '%s' created.", htmlEscape(fileTarget))
	if _, err := os.Stat(fileTarget); os.IsNotExist(err) {
		fmt.Fprintln(w, " But file doesn't exist yet.")
	} else {
		fmt.Fprintln(w, " And file already exists.")
	}
}

func doSomething(param string) string {
	bar := "safe!"
	map35951 := map[string]interface{}{
		"keyA-35951": "a-Value",
		"keyB-35951": param,
		"keyC":       "another-Value",
	}
	bar = map35951["keyB-35951"].(string)
	return bar
}

func htmlEscape(s string) string {
	return template.HTMLEscapeString(s)
}

func main() {
	http.Handle("/pathtraver-02/BenchmarkTest02197", &BenchmarkTest02197{})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
