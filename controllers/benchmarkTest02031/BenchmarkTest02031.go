package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func main() {
	http.HandleFunc("/pathtraver-02/BenchmarkTest02031", BenchmarkTest02031)
	http.ListenAndServe(":8080", nil)
}

func BenchmarkTest02031(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := r.Header.Get("BenchmarkTest02031")
	param, _ = url.QueryUnescape(param)

	bar := doSomething(r, param)

	fileName := fmt.Sprintf("testfiles/%s", bar)
	fis, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintf(w, "Couldn't open FileInputStream on file: '%s'", fileName)
		fmt.Fprintf(w, "Problem getting FileInputStream: %s", err.Error())
		return
	}
	defer fis.Close()

	b, err := ioutil.ReadAll(fis)
	if err != nil {
		fmt.Fprintf(w, "Error reading file: %s", err.Error())
		return
	}

	_, _ = w.Write([]byte(fmt.Sprintf("The beginning of file: '%s' is:\n\n%s", fileName, b)))
}

func doSomething(r *http.Request, param string) string {
	num := 106
	if (7*18)+num > 200 {
		return "This_should_always_happen"
	}
	return param
}
