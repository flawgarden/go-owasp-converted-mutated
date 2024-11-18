package controllers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

type BenchmarkTest01409 struct{}

func (b *BenchmarkTest01409) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		b.doPost(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (b *BenchmarkTest01409) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	flag := true
	for name, values := range r.Form {
		if !flag {
			break
		}
		for _, value := range values {
			if value == "BenchmarkTest01409" {
				param = name
				flag = false
				break
			}
		}
	}

	bar := b.doSomething(param)

	fileName := filepath.Join("testfiles", bar)
	file, err := os.Open(fileName)
	if err != nil {
		http.Error(w, fmt.Sprintf("Couldn't open InputStream on file: %s", fileName), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	content := make([]byte, 1000)
	size, err := file.Read(content)
	if err != nil {
		http.Error(w, fmt.Sprintf("Problem getting InputStream: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(fmt.Sprintf("The beginning of file: '%s' is:\n\n", fileName)))
	w.Write(content[:size])
}

func (b *BenchmarkTest01409) doSomething(param string) string {
	bar := "safe!"
	map58007 := make(map[string]interface{})
	map58007["keyA-58007"] = "a_Value"
	map58007["keyB-58007"] = param
	map58007["keyC"] = "another_Value"
	bar = map58007["keyB-58007"].(string)
	bar = map58007["keyA-58007"].(string)

	return bar
}

func main() {
	http.Handle("/pathtraver-01/BenchmarkTest01409", &BenchmarkTest01409{})
	http.ListenAndServe(":8080", nil)
}
