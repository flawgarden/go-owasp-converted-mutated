package controllers

import (
	"fmt"
	"net/http"
	"os"
)

type BenchmarkTest00215 struct{}

func (b *BenchmarkTest00215) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	var param string
	for name := range r.Header {
		if !isCommonHeader(name) {
			param = name
			break
		}
	}

	bar := param
	fileTarget := fmt.Sprintf("%s/%s", os.Getenv("TESTFILES_DIR"), bar)

	w.Write([]byte(fmt.Sprintf("Access to file: '%s' created.", htmlEncode(fileTarget))))
	if _, err := os.Stat(fileTarget); os.IsNotExist(err) {
		w.Write([]byte(" But file doesn't exist yet."))
	} else {
		w.Write([]byte(" And file already exists."))
	}
}

func isCommonHeader(header string) bool {
	// Define common headers here
	return false // Placeholder implementation
}

func htmlEncode(s string) string {
	// Simple HTML encoding implementation
	return s // Placeholder implementation
}

func main() {
	http.Handle("/pathtraver-00/BenchmarkTest00215", &BenchmarkTest00215{})
	http.ListenAndServe(":8080", nil)
}
