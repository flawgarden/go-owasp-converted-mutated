package controllers

import (
	"fmt"
	"html"
	"net/http"
	"os"
	"path/filepath"
)

type BenchmarkTest00457 struct{}

func (b *BenchmarkTest00457) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		b.doPost(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (b *BenchmarkTest00457) doPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest00457")

	bar := ""
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the param value

var i123 interface{} = &Anon{Value1: bar}
if ptr, ok := i123.(*EmbeddedStruct); ok {
     bar = ptr.Field1
} else {
    bar = "YGYXI"
}

	}

	fileName := filepath.Join("testfiles", bar)
	fos, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Couldn't open FileOutputStream on file: '" + fileName + "'")
		return
	}
	defer fos.Close()

	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	w.Write([]byte("Now ready to write to file: " + html.EscapeString(fileName)))
}
