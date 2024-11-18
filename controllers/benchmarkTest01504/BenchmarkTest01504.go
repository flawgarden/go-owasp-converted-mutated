package controllers

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"html"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01504 struct{}

func (b *BenchmarkTest01504) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest01504")
	if param == "" {
		param = ""
	}

	bar := b.doSomething(param)

	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	md := md5.New()
	_, err := md.Write([]byte(bar))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	result := md.Sum(nil)

	fileTarget := "passwordFile.txt"
	fw, err := os.OpenFile(fileTarget, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer fw.Close()
	if _, err := fw.WriteString(fmt.Sprintf("hash_value=%s\n", base64.StdEncoding.EncodeToString(result))); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Sensitive value '%s' hashed and stored<br/>", html.EscapeString(bar))
}

func (b *BenchmarkTest01504) doSomething(param string) string {
	num := 106
	bar := "This should never happen"
	if (7*42)-num <= 200 {
		bar = param
	}
	return bar
}

func main() {
	http.Handle("/hash-01/BenchmarkTest01504", &BenchmarkTest01504{})
	http.ListenAndServe(":8080", nil)
}
