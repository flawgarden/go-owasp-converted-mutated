package controllers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

type BenchmarkTest02105 struct{}

func (b *BenchmarkTest02105) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.URL.Query().Get("BenchmarkTest02105")
	if param == "" {
		param = ""
	}

	bar := doSomething(param)

	fileTarget := filepath.Join("testfiles", bar)
	fmt.Fprintf(w, "Access to file: '%s' created.\n", fileTarget)
	if _, err := os.Stat(fileTarget); os.IsNotExist(err) {
		fmt.Fprintf(w, "But file doesn't exist yet.\n")
	} else {
		fmt.Fprintf(w, "And file already exists.\n")
	}
}

func doSomething(param string) string {
	bar := ""
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = append(valuesList[:0], valuesList[1:]...) // remove the 1st safe value
		bar = valuesList[0]                                    // get the param value
	}
	return bar
}

func main() {
	http.Handle("/pathtraver-02/BenchmarkTest02105", &BenchmarkTest02105{})
	http.ListenAndServe(":8080", nil)
}
