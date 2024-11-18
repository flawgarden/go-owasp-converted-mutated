package controllers

import (
	"net/http"
	"os"
)

type BenchmarkTest01113 struct{}

func (b *BenchmarkTest01113) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	var param string
	for name := range r.Header {
		if isCommonHeader(name) {
			continue
		}
		param = name
		break
	}

	bar := new(Test).doSomething(r, param)

	fileName := "path/to/directory/" + bar
	fos, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		http.Error(w, "Couldn't open FileOutputStream", http.StatusInternalServerError)
		return
	}
	defer fos.Close()

	w.Write([]byte("Now ready to write to file: " + fileName))
}

func isCommonHeader(name string) bool {
	// Implement logic to check common headers
	return false
}

type Test struct{}

func (t *Test) doSomething(r *http.Request, param string) string {
	bar := "safe!"
	map50384 := make(map[string]interface{})
	map50384["keyA-50384"] = "a_Value"
	map50384["keyB-50384"] = param
	map50384["keyC"] = "another_Value"
	bar = map50384["keyB-50384"].(string)
	bar = map50384["keyA-50384"].(string)

	return bar
}
