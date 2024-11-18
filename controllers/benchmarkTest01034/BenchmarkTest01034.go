package controllers

import (
	"encoding/base64"
	"html"
	"net/http"
	"os"
)

type BenchmarkTest01034Controller struct{}

func (c *BenchmarkTest01034Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	if r.Header.Get("BenchmarkTest01034") != "" {
		param = r.Header.Get("BenchmarkTest01034")
	}

	bar := new(Test).doSomething(r, param)

	fileName := ""
	var fos *os.File

	defer func() {
		if fos != nil {
			fos.Close()
		}
	}()

	fileName = "path/to/testfiles/" + bar

	fos, err := os.Create(fileName)
	if err != nil {
		http.Error(w, "Couldn't open FileOutputStream", http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Now ready to write to file: " + escapeHTML(fileName)))
}

type Test struct{}

func (t *Test) doSomething(r *http.Request, param string) string {
	bar := ""
	if param != "" {
		decoded, _ := base64.StdEncoding.DecodeString(param)
		bar = string(decoded)
	}
	return bar
}

func escapeHTML(s string) string {
	return html.EscapeString(s)
}
