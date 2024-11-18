package controllers

import (
	"fmt"
	"net/http"
	"os"
)

type BenchmarkTest02202 struct{}

func (b *BenchmarkTest02202) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest02202")

	bar := doSomething(param)

	fileName := ""
	var fos *os.File

	defer func() {
		if fos != nil {
			fos.Close()
		}
	}()

	fileName = "path/to/testfiles/" + bar

	var err error
	fos, err = os.Create(fileName)
	if err != nil {
		fmt.Println("Couldn't open FileOutputStream on file: '" + fileName + "'")
		return
	}

	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	w.Write([]byte("Now ready to write to file: " + fileName))
}

func doSomething(param string) string {
	bar := ""
	num := 86
	if (7*42)-num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}
	return bar
}

func main() {
	http.Handle("/pathtraver-02/BenchmarkTest02202", &BenchmarkTest02202{})
	http.ListenAndServe(":8080", nil)
}
