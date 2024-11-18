package controllers

import (
	"fmt"
	"net/http"
	"os"
)

type BenchmarkTest00863Controller struct {
	http.ResponseWriter
	*http.Request
}

func (c *BenchmarkTest00863Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest00863")
	bar := doSomething(param)

	fileTarget := fmt.Sprintf("%s/Test.txt", bar)
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	_, err := w.Write([]byte(fmt.Sprintf("Access to file: '%s' created.", fileTarget)))
	if err != nil {
		panic(err)
	}
	if _, err := os.Stat(fileTarget); os.IsNotExist(err) {
		w.Write([]byte(" But file doesn't exist yet."))
	} else {
		w.Write([]byte(" And file already exists."))
	}
}

func doSomething(param string) string {
	// Здесь должна быть реализация создания объекта и выполнение необходимых операций
	return param
}
