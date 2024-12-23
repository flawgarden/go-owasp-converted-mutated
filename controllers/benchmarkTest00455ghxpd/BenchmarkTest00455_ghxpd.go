package controllers

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const testFilesDir = "/path/to/test/files/"

type BenchmarkTest00455 struct{}

func (b *BenchmarkTest00455) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest00455")

	bar := ""
	if param != "" {
		decoded, _ := base64.StdEncoding.DecodeString(param)
		bar = string(decoded)

nested7231 := NewNestedFields4(bar)
bar = nested7231.nested1.nested1.nested1.nested1.value

	}

	fileName := testFilesDir + bar
	fis, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintf(w, "Couldn't open FileInputStream on file: '%s'", fileName)
		return
	}
	defer fis.Close()

	binaryData, _ := ioutil.ReadAll(fis)
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	fmt.Fprintf(w, "The beginning of file: '%s' is:\n\n", fileName)
	w.Write(binaryData[:1000])
}

func main() {
	http.Handle("/pathtraver-00/BenchmarkTest00455", &BenchmarkTest00455{})
	http.ListenAndServe(":8080", nil)
}
