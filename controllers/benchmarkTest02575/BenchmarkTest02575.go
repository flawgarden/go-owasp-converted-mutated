package controllers

import (
	"crypto/md5"
	"encoding/base64"
	"html"
	"net/http"
	"net/url"
	"os"
)

type BenchmarkTest02575 struct{}

func (b *BenchmarkTest02575) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		b.doPost(w, r)
		return
	}
	b.doPost(w, r)
}

func (b *BenchmarkTest02575) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	queryString := r.URL.RawQuery
	paramval := "BenchmarkTest02575="
	paramLoc := -1
	if queryString != "" {
		paramLoc = len(queryString) - len(paramval)
		if paramLoc == -1 {
			http.Error(w, "getQueryString() couldn't find expected parameter 'BenchmarkTest02575' in query string.", http.StatusBadRequest)
			return
		}
	}

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := len(queryString)
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : ampersandLoc]
	}
	param, _ = url.QueryUnescape(param)

	bar := doSomething(param)

	hash := md5.Sum([]byte(bar))
	fileTarget, _ := os.OpenFile("passwordFile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer fileTarget.Close()
	fileTarget.WriteString("hash_value=" + base64.StdEncoding.EncodeToString(hash[:]) + "\n")

	w.Write([]byte("Sensitive value '" + html.EscapeString(bar) + "' hashed and stored<br/>"))
	w.Write([]byte("Hash Test executed"))
}

func doSomething(param string) string {
	bar := param
	num := 86
	if (7*42)-num > 200 {
		bar = "This_should_always_happen"
	}
	return bar
}
