package controllers

import (
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type BenchmarkTest00786 struct{}

func (b *BenchmarkTest00786) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	queryString := r.URL.RawQuery
	paramval := "BenchmarkTest00786="
	paramLoc := strings.Index(queryString, paramval)
	if paramLoc == -1 {
		http.Error(w, "getQueryString() couldn't find expected parameter 'BenchmarkTest00786' in query string.", http.StatusBadRequest)
		return
	}

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := strings.Index(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : paramLoc+ampersandLoc]
	}
	param, _ = url.QueryUnescape(param)

	bar := "safe!"
	map29957 := make(map[string]string)
	map29957["keyA-29957"] = "a_Value"
	map29957["keyB-29957"] = param
	map29957["keyC"] = "another_Value"
	bar = map29957["keyB-29957"]

	fileName := bar
	fos, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		http.Error(w, "Couldn't open FileOutputStream on file: '"+fileName+"'", http.StatusInternalServerError)
		return
	}
	defer fos.Close()

	response := map[string]string{"message": "Now ready to write to file: " + fileName}
	output, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func main() {
	http.Handle("/pathtraver-00/BenchmarkTest00786", &BenchmarkTest00786{})
	http.ListenAndServe(":8080", nil)
}
