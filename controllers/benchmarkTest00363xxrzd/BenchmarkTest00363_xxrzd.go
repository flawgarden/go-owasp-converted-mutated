package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type BenchmarkTest00363 struct{}

func (b BenchmarkTest00363) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := r.URL.Query().Get("BenchmarkTest00363")

var a12341 BinaryOpInterface = &ImplBinaryOpInterfaceClass2{}
param = a12341.InterfaceCall(param, "")

	if param == "" {
		param = ""
	}
	bar := param
	fileName := "testfiles/" + bar

	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Fprintln(w, "Couldn't open FileOutputStream on file: '"+fileName+"'")
		return
	}
	defer f.Close()

	response := map[string]string{
		"message": "Now ready to write to file: " + escapeHTML(fileName),
	}
	output, _ := json.Marshal(response)
	w.Write(output)
}

func escapeHTML(input string) string {
	return input // реализация для экранирования HTML
}

func main() {
	http.Handle("/pathtraver-00/BenchmarkTest00363", BenchmarkTest00363{})
	http.ListenAndServe(":8080", nil)
}
