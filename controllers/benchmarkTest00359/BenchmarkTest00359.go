package controllers

import (
	"fmt"
	"net/http"
	"os"
)

type BenchmarkTest00359 struct{}

func (b *BenchmarkTest00359) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.URL.Query().Get("BenchmarkTest00359")
	if param == "" {
		param = ""
	}

	var bar string
	num := 196
	if (500/42)+num > 200 {
		bar = param
	} else {
		bar = "This should never happen"
	}

	fileTarget := bar
	w.Write([]byte(fmt.Sprintf("Access to file: '%s' created.", fileTarget)))
	if _, err := os.Stat(fileTarget); err == nil {
		w.Write([]byte(" And file already exists."))
	} else {
		w.Write([]byte(" But file doesn't exist yet."))
	}
}

func main() {
	http.Handle("/pathtraver-00/BenchmarkTest00359", &BenchmarkTest00359{})
	http.ListenAndServe(":8080", nil)
}
