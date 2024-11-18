package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"
	"strings"
)

type BenchmarkTest00794 struct{}

func (b *BenchmarkTest00794) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "This input source requires a POST, not a GET.", http.StatusMethodNotAllowed)
		return
	}

	queryString := r.URL.RawQuery
	paramval := "BenchmarkTest00794="
	paramLoc := strings.Index(queryString, paramval)
	if paramLoc == -1 {
		http.Error(w, "getQueryString() couldn't find expected parameter 'BenchmarkTest00794' in query string.", http.StatusBadRequest)
		return
	}

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := strings.Index(param, "&")
	if ampersandLoc != -1 {
		param = param[:ampersandLoc]
	}

	bar := param
	if len(param) > 1 {
		bar = param[:len(param)-1]
	}

	hash := md5.Sum([]byte(bar))
	hashValue := hex.EncodeToString(hash[:])

	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	w.Write([]byte("Sensitive value '" + bar + "' hashed and stored<br/>"))
	w.Write([]byte("Hash Test executed, hash value: " + hashValue))
}
