package controllers

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
)

type BenchmarkTest01761 struct{}

func (b *BenchmarkTest01761) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		b.doPost(w, r)
	} else if r.Method == http.MethodPost {
		b.doPost(w, r)
	}
}

func (b *BenchmarkTest01761) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.URL.Query().Get("BenchmarkTest01761")

	bar := b.doSomething(param)

	hash := sha1.New()
	input := []byte(bar)
	hash.Write(input)
	result := hash.Sum(nil)

	fileTarget, err := os.OpenFile("passwordFile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		http.Error(w, "Unable to open file", http.StatusInternalServerError)
		return
	}
	defer fileTarget.Close()

	_, err = fileTarget.WriteString("hash_value=" + base64.StdEncoding.EncodeToString(result) + "\n")
	if err != nil {
		http.Error(w, "Unable to write to file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Sensitive value '%s' hashed and stored<br/>", htmlEscape(string(input)))
	fmt.Fprintf(w, "Hash Test executed")
}

func (b *BenchmarkTest01761) doSomething(param string) string {
	thing := createThing()
	return thing.doSomething(param)
}

func createThing() ThingInterface {
	return &Thing{}
}

type ThingInterface interface {
	doSomething(param string) string
}

type Thing struct{}

func (t *Thing) doSomething(param string) string {
	return param
}

func htmlEscape(s string) string {
	return s // Вставьте реальную реализацию HTML экранирования
}
