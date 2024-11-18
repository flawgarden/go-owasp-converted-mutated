package controllers

import (
	"fmt"
	"net/http"
	"strings"
)

type BenchmarkTest01407 struct{}

func (b *BenchmarkTest01407) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		b.doPost(w, r)
	} else if r.Method == http.MethodPost {
		b.doPost(w, r)
	}
}

func (b *BenchmarkTest01407) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	flag := true
	names := r.URL.Query()

	for name, values := range names {
		for _, value := range values {
			if value == "BenchmarkTest01407" {
				param = name
				flag = false
				break
			}
		}
		if !flag {
			break
		}
	}

	bar := b.doSomething(param)

	fileName := "path/to/testfiles/" + bar

	err := writeFile(fileName)
	if err != nil {
		fmt.Println("Couldn't open FileOutputStream on file: '" + fileName + "'")
	}

	_, _ = w.Write([]byte("Now ready to write to file: " + escapeHTML(fileName)))
}

func (b *BenchmarkTest01407) doSomething(param string) string {
	bar := ""

	num := 86
	if (7*42)-num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}

	return bar
}

func writeFile(fileName string) error {
	// Dummy implementation for file writing
	return nil
}

func escapeHTML(s string) string {
	return strings.ReplaceAll(s, "&", "&amp;")
}

func main() {
	http.Handle("/pathtraver-01/BenchmarkTest01407", &BenchmarkTest01407{})
	http.ListenAndServe(":8080", nil)
}
