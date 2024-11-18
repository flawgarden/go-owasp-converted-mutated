package controllers

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"html"
	"net/http"
	"net/url"
	"os"
)

type BenchmarkTest00270 struct {
	http.Handler
}

func (b *BenchmarkTest00270) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	var param string
	headers := r.Header["BenchmarkTest00270"]
	if len(headers) > 0 {
		param = headers[0]
	}

	param, _ = url.QueryUnescape(param)

	var bar string
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:]
		bar = valuesList[0]
	}

	md := sha512.New()
	input := []byte("?")
	if bar != "" {
		input = []byte(bar)
	}
	md.Write(input)

	result := md.Sum(nil)
	fileTarget := "passwordFile.txt"
	fw, err := os.OpenFile(fileTarget, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(w, "Error opening file", http.StatusInternalServerError)
		return
	}
	defer fw.Close()
	fw.WriteString(fmt.Sprintf("hash_value=%s\n", base64.StdEncoding.EncodeToString(result)))

	w.Write([]byte(fmt.Sprintf("Sensitive value '%s' hashed and stored<br/>", html.EscapeString(string(input)))))
	w.Write([]byte("Hash Test executed"))
}
