package controllers

import (
	"fmt"
	"html"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

type BenchmarkTest01035 struct{}

func (b *BenchmarkTest01035) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.Header.Get("BenchmarkTest01035")
	param, _ = url.QueryUnescape(param)

	bar := b.doSomething(r, param)

	fileName := ""
	var fos *os.File

	defer func() {
		if fos != nil {
			fos.Close()
		}
	}()

	fileName = filepath.Join("testfiles", bar)

	fos, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Couldn't open FileOutputStream on file: '" + fileName + "'")
		return
	}

	w.Write([]byte("Now ready to write to file: " + html.EscapeString(fileName)))
}

func (b *BenchmarkTest01035) doSomething(r *http.Request, param string) string {
	var bar string

	num := 86
	if (7*42)-num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}

	return bar
}

func main() {
	http.Handle("/pathtraver-01/BenchmarkTest01035", &BenchmarkTest01035{})
	http.ListenAndServe(":8080", nil)
}
