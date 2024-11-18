package controllers

import (
	"fmt"
	"net/http"
	"os"
)

type BenchmarkTest01239 struct{}

func (b *BenchmarkTest01239) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.URL.Query().Get("BenchmarkTest01239")
	if param == "" {
		param = ""
	}

	bar := new(Test).doSomething(r, param)

	fileName := "./testfiles/" + bar

	fos, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("Couldn't open FileOutputStream on file: '" + fileName + "'")
		return
	}
	defer fos.Close()

	_, err = w.Write([]byte("Now ready to write to file: " + fileName))
	if err != nil {
		fmt.Println("Error writing response")
	}
}

type Test struct{}

func (t *Test) doSomething(r *http.Request, param string) string {
	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}

		valuesList = valuesList[1:] // remove the 1st safe value

		bar = valuesList[1] // get the last 'safe' value
	}

	return bar
}
