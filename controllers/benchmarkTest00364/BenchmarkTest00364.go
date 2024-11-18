package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

type BenchmarkTest00364 struct{}

func (b *BenchmarkTest00364) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	response := "text/html;charset=UTF-8"
	w.Header().Set("Content-Type", response)

	param := r.FormValue("BenchmarkTest00364")
	if param == "" {
		param = ""
	}

	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[1]         // get the last 'safe' value
	}

	fileName := fmt.Sprintf("testfiles/%s", bar)
	fos, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		http.Error(w, fmt.Sprintf("Couldn't open FileOutputStream on file: '%s'", fileName), http.StatusInternalServerError)
		return
	}
	defer fos.Close()

	output := fmt.Sprintf("Now ready to write to file: %s", htmlEscape(fileName))
	w.Write([]byte(output))
}

func htmlEscape(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(s, "&", "&amp;"), "<", "&lt;")
}

func main() {
	http.Handle("/pathtraver-00/BenchmarkTest00364", &BenchmarkTest00364{})
	http.ListenAndServe(":8080", nil)
}
