package controllers

import (
	"net/http"
	"os"
	"path/filepath"
)

type BenchmarkTest01231 struct{}

func (b *BenchmarkTest01231) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		b.doPost(w, r)
	} else if r.Method == http.MethodPost {
		b.doPost(w, r)
	}
}

func (b *BenchmarkTest01231) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := r.URL.Query().Get("BenchmarkTest01231")
	if param == "" {
		param = ""
	}

	bar := new(Test).doSomething(r, param)

	fileTarget := filepath.Join(os.Getenv("TESTFILES_DIR"), bar)
	w.Write([]byte("Access to file: '" + fileTarget + "' created."))
	if _, err := os.Stat(fileTarget); os.IsNotExist(err) {
		w.Write([]byte(" But file doesn't exist yet."))
	} else {
		w.Write([]byte(" And file already exists."))
	}
}

type Test struct{}

func (t *Test) doSomething(r *http.Request, param string) string {
	bar := "safe!"
	map35717 := make(map[string]interface{})
	map35717["keyA-35717"] = "a-Value"
	map35717["keyB-35717"] = param
	map35717["keyC"] = "another-Value"
	bar = map35717["keyB-35717"].(string)
	return bar
}
