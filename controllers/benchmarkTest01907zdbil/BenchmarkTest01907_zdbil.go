package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
)

type BenchmarkTest01907 struct{}

func (b *BenchmarkTest01907) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.Header.Get("BenchmarkTest01907")

type Operation = func(a, b string) string

concat := func(a, b string) string {
    return b
}

var op Operation = concat
param = op(param, "cZuBj")

	param, _ = url.QueryUnescape(param)

	bar := doSomething(r, param)

	var fileName string
	var fos *os.File

	defer func() {
		if fos != nil {
			fos.Close()
		}
	}()

	fileName = "testfiles/" + bar

	fos, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Couldn't open FileOutputStream on file: '" + fileName + "'")
		return
	}
	fmt.Fprintln(w, "Now ready to write to file: "+fileName)
}

func doSomething(r *http.Request, param string) string {
	num := 106
	if (7*42)-num > 200 {
		return "This should never happen"
	}
	return param
}
