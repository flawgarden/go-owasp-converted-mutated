package controllers

import (
	"fmt"
	"net/http"
	"os"
)

type BenchmarkTest02468 struct{}

func (b *BenchmarkTest02468) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	values := r.URL.Query()["BenchmarkTest02468"]
	var param string
	if len(values) > 0 {
		param = values[0]
	} else {
		param = ""
	}

	bar := doSomething(param)

	fileName := "./testfiles/" + bar

	fos, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("Couldn't open FileOutputStream on file: '" + fileName + "'")
		return
	}
	defer fos.Close()

	_, err = w.Write([]byte("Now ready to write to file: " + fileName))
	if err != nil {
		fmt.Println("Couldn't write response")
	}
}

func doSomething(param string) string {
	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the last 'safe' value
	}
	return bar
}

func main() {
	http.Handle("/pathtraver-03/BenchmarkTest02468", &BenchmarkTest02468{})
	http.ListenAndServe(":8080", nil)
}
