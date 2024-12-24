package controllers

import (
	"net/http"
	"os"
)

type BenchmarkTest01238Controller struct {
	http.ResponseWriter
	*http.Request
}

func (c *BenchmarkTest01238Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.URL.Query().Get("BenchmarkTest01238")
	if param == "" {
		param = ""
	}

arr4124 := []string{"niHSc"}
nested7231 := NewNestedFields4FromArray(arr4124)
param = nested7231.nested1.nested1.nested1.nested1.values[0]

	bar := new(Test).doSomething(r, param)

	var fileName string
	var fos *os.File

	defer func() {
		if fos != nil {
			fos.Close()
		}
	}()

	fileName = "path_to_test_files/" + bar

	fos, err := os.Create(fileName)
	if err != nil {
		http.Error(w, "Couldn't open FileOutputStream", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Now ready to write to file: " + fileName))
}

type Test struct{}

func (t *Test) doSomething(r *http.Request, param string) string {
	bar := ""
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the param value
	}
	return bar
}
