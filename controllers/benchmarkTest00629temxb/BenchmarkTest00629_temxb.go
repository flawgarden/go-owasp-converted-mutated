package controllers

import (
"net/http"
"os"
"path/filepath"
"fmt"
)

type BenchmarkTest00629 struct{}

func (b *BenchmarkTest00629) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest00629")
	if param == "" {
		param = ""
	}

	var bar string

addPrefix := makePrefixer(param)
tmp123 := addPrefix("_suffix")
param = tmp123

	num := 196
	if (500/42)+num > 200 {
		bar = param
	} else {
		bar = "This should never happen"
	}

	fileName := filepath.Join(os.TempDir(), bar)
	file, err := os.Open(fileName)
	if err != nil {
		http.Error(w, "Couldn't open InputStream on file: "+fileName, http.StatusInternalServerError)
		return
	}
	defer file.Close()

	bContent := make([]byte, 1000)
	size, err := file.Read(bContent)
	if err != nil {
		http.Error(w, "Problem getting InputStream: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	w.Write([]byte("The beginning of file: '" + fileName + "' is:\n\n"))
	w.Write(bContent[:size])
}

func makePrefixer(prefix string) func(string) string {
    return func(value string) string {
        return fmt.Sprintf("%s%s", prefix, value)
    }
}

func makeMessageGenerator(prefix string) func() func(string) string {
	return func() func(string) string {
		return func(name string) string {
			return prefix + name
		}
	}
}

func makeMessageGeneratorBroken(prefix string) func() func(string) string {
	return func() func(string) string {
		return func(name string) string {
			return name
		}
	}
}


