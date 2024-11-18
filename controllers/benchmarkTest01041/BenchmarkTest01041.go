package controllers

import (
	"crypto/sha512"
	"encoding/base64"
	"hash"
	"html"
	"net/http"
	"net/url"
	"os"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01041 struct {
}

func (b *BenchmarkTest01041) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.Header.Get("BenchmarkTest01041")
	param, _ = url.QueryUnescape(param)

	bar := b.doSomething(param)

	algorithm := "SHA512"
	hasher := GetDigest(algorithm)
	input := []byte(bar)
	hasher.Write(input)

	result := hasher.Sum(nil)
	fileTarget := "passwordFile.txt"
	fw, err := os.OpenFile(fileTarget, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(w, "Could not open file", http.StatusInternalServerError)
		return
	}
	defer fw.Close()

	_, err = fw.WriteString("hash_value=" + EncodeForBase64(result) + "\n")
	if err != nil {
		http.Error(w, "Could not write to file", http.StatusInternalServerError)
		return
	}

	_, _ = w.Write([]byte("Sensitive value '" + EscapeHTML(string(input)) + "' hashed and stored<br/>"))
	_, _ = w.Write([]byte("Hash Test java.security.MessageDigest.getInstance(java.lang.String) executed"))
}

func (b *BenchmarkTest01041) doSomething(param string) string {
	bar := ""
	if param != "" {
		bar = param
	}
	return bar
}

func GetDigest(algorithm string) hash.Hash {
	switch algorithm {
	case "SHA512":
		h := sha512.New()
		return h
	}
	return nil
}

func EncodeForBase64(input []byte) string {
	return base64.StdEncoding.EncodeToString(input)
}

func EscapeHTML(input string) string {
	return html.EscapeString(input)
}
