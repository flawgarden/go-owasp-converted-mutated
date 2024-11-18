package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
)

type BenchmarkTest00134 struct{}

func (b *BenchmarkTest00134) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.Header.Get("BenchmarkTest00134")
	param, _ = url.QueryUnescape(param)

	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the last 'safe' value
	}

	fileName := "/path/to/testfiles/" + bar

	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Couldn't open FileOutputStream on file:", fileName)
		return
	}
	defer file.Close()

	_, err = w.Write([]byte("Now ready to write to file: " + fileName))
	if err != nil {
		fmt.Println("Error writing response:", err)
	}
}

func main() {
	http.Handle("/pathtraver-00/BenchmarkTest00134", &BenchmarkTest00134{})
	http.ListenAndServe(":8080", nil)
}
