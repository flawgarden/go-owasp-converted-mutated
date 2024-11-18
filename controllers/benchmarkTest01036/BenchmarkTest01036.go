package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

type BenchmarkTest01036 struct{}

func (b *BenchmarkTest01036) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		b.doPost(w, r)
	} else if r.Method == http.MethodPost {
		b.doPost(w, r)
	}
}

func (b *BenchmarkTest01036) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.Header.Get("BenchmarkTest01036")
	param, _ = url.QueryUnescape(param)

	bar := b.doSomething(r, param)

	fileName := filepath.Join("testfiles", bar)
	fileData, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Fprintf(w, "Couldn't open InputStream on file: '%s'", fileName)
		fmt.Fprintf(w, "Problem getting InputStream: %s", err.Error())
		return
	}

	fmt.Fprintf(w, "The beginning of file: '%s' is:\n\n", sanitize(fileName))
	fmt.Fprintf(w, sanitize(string(fileData[:1000])))
}

func (b *BenchmarkTest01036) doSomething(r *http.Request, param string) string {
	bar := "safe!"
	map8487 := make(map[string]interface{})
	map8487["keyA-8487"] = "a_Value"
	map8487["keyB-8487"] = param
	map8487["keyC"] = "another_Value"
	bar = map8487["keyB-8487"].(string)
	bar = map8487["keyA-8487"].(string)

	return bar
}

func sanitize(input string) string {
	return input // Пример простой санитации
}

func main() {
	http.Handle("/pathtraver-01/BenchmarkTest01036", &BenchmarkTest01036{})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server failed to start:", err)
		os.Exit(1)
	}
}
