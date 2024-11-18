package controllers

import (
	"net/http"
	"os"
)

type BenchmarkTest01751 struct{}

func (bt *BenchmarkTest01751) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest01751")

	bar := bt.doSomething(r, param)

	fileName := ""
	var fos *os.File

	defer func() {
		if fos != nil {
			fos.Close()
		}
	}()

	fileName = "path/to/directory/" + bar // Adjust directory as necessary

	fos, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		http.Error(w, "Couldn't open file: "+fileName, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	w.Write([]byte("Now ready to write to file: " + fileName))
}

func (bt *BenchmarkTest01751) doSomething(r *http.Request, param string) string {
	var bar string
	num := 106
	if 7*18+num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}
	return bar
}
