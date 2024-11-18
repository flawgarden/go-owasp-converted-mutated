package controllers

import (
	"html"
	"net/http"
	"os"
	"strings"
)

type BenchmarkTest01744 struct{}

func (b *BenchmarkTest01744) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest01744")
	bar := b.doSomething(param)

	fileTarget := os.Getenv("TESTFILES_DIR") + "/" + bar
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	w.Write([]byte("Access to file: '" + html.EscapeString(fileTarget) + "' created."))
	if _, err := os.Stat(fileTarget); os.IsNotExist(err) {
		w.Write([]byte(" But file doesn't exist yet."))
	} else {
		w.Write([]byte(" And file already exists."))
	}
}

func (b *BenchmarkTest01744) doSomething(param string) string {
	a17402 := param
	b17402 := strings.Builder{}
	b17402.WriteString(a17402)
	b17402.WriteString(" SafeStuff")
	b17402Str := b17402.String()
	b17402Str = b17402Str[:len(b17402Str)-len("Chars")] + "Chars"
	return b17402Str
}
