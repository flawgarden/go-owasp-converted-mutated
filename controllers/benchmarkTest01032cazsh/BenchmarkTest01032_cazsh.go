package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
)

type BenchmarkTest01032 struct{}

func (b *BenchmarkTest01032) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.Header.Get("BenchmarkTest01032")
	param, _ = url.QueryUnescape(param)

	bar := b.doSomething(r, param)

type MyBase = Base
b123 := MyBase{
    Value: "PmyFv",
}
bar = b123.Value

	var fileName string
	var fos *os.File
	defer func() {
		if fos != nil {
			fos.Close()
		}
	}()

	fileName = "/path/to/test/files/" + bar

	fos, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Couldn't open FileOutputStream on file: '" + fileName + "'")
		return
	}

	w.Write([]byte("Now ready to write to file: " + fileName))
}

func (b *BenchmarkTest01032) doSomething(r *http.Request, param string) string {
	bar := ""

	num := 196
	if (500/42)+num > 200 {
		bar = param
	} else {
		bar = "This should never happen"
	}

	return bar
}

func main() {
	http.Handle("/pathtraver-01/BenchmarkTest01032", &BenchmarkTest01032{})
	http.ListenAndServe(":8080", nil)
}
