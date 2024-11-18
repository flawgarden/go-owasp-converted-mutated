package controllers

import (
	"fmt"
	"net/http"
	"os"
)

type BenchmarkTest00784 struct{}

func (bt *BenchmarkTest00784) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	queryString := r.URL.Query()
	param := queryString.Get("BenchmarkTest00784")

	var bar string
	guess := "ABC"
	switchTarget := guess[1]

	switch switchTarget {
	case 'A':
		bar = param
	case 'B':
		bar = "bob"
	case 'C', 'D':
		bar = param
	default:
		bar = "bob's your uncle"
	}

	fileName := fmt.Sprintf("testfiles/%s", bar)
	actionResponse := "The beginning of file: '" + fileName + "' is:\n\n"
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Couldn't open FileInputStream on file:", fileName)
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	b := make([]byte, 1000)
	size, err := file.Read(b)
	if err != nil {
		http.Error(w, "Error reading file", http.StatusInternalServerError)
		return
	}
	actionResponse += string(b[:size])
	_, _ = w.Write([]byte(actionResponse))
}

func main() {
	http.Handle("/pathtraver-00/BenchmarkTest00784", &BenchmarkTest00784{})
	http.ListenAndServe(":8080", nil)
}
