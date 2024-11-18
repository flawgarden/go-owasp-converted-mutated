package controllers

import (
	"fmt"
	"net/http"
	"os"
)

type BenchmarkTest struct{}

func (b *BenchmarkTest) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest02377")
	if param == "" {
		param = ""
	}

	bar := b.doSomething(param)

	fileTarget := fmt.Sprintf("%s/Test.txt", bar)
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	fmt.Fprintf(w, "Access to file: '%s' created.", htmlEscape(fileTarget))
	if _, err := os.Stat(fileTarget); os.IsNotExist(err) {
		fmt.Fprint(w, " But file doesn't exist yet.")
	} else {
		fmt.Fprint(w, " And file already exists.")
	}
}

func (b *BenchmarkTest) doSomething(param string) string {
	bar := param
	return bar
}

func htmlEscape(str string) string {
	// Simple HTML escaping for demonstration purposes
	return jsonEscape(str)
}

func jsonEscape(str string) string {
	return fmt.Sprintf("%q", str)
}

func main() {
	http.Handle("/pathtraver-02/BenchmarkTest02377", &BenchmarkTest{})
	http.ListenAndServe(":8080", nil)
}
