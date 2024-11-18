package controllers

import (
	"net/http"
	"os"
)

type BenchmarkTest01330 struct{}

func (b *BenchmarkTest01330) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.URL.Query().Get("BenchmarkTest01330")

	bar := new(Test).doSomething(r, param)

	fileName := ""
	var fos *os.File

	defer func() {
		if fos != nil {
			fos.Close()
		}
	}()

	fileName = "./testfiles/" + bar

	fos, err := os.Create(fileName)
	if err != nil {
		http.Error(w, "Couldn't open FileOutputStream on file: "+fileName, http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Now ready to write to file: " + fileName))
}

type Test struct{}

func (t *Test) doSomething(r *http.Request, param string) string {
	bar := ""

	num := 196
	if (500/42)+num > 200 {
		bar = param
	} else {
		bar = "This should never happen"
	}

	return bar
}
