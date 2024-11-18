package controllers

import (
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"html"
	"net/http"
	"os"
)

type BenchmarkTest02216 struct{}

func (bt *BenchmarkTest02216) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	if r.Method == http.MethodGet {
		bt.doPost(w, r)
		return
	}
	if r.Method == http.MethodPost {
		bt.doPost(w, r)
		return
	}
	http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
}

func (bt *BenchmarkTest02216) doPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest02216")
	bar := doSomething(param)

	hasher := sha512.New()
	input := []byte(bar)
	hasher.Write(input)
	result := hasher.Sum(nil)

	file, err := os.OpenFile("passwordFile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(w, "Unable to open file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	if _, err := file.WriteString("hash_value=" + base64.StdEncoding.EncodeToString(result) + "\n"); err != nil {
		http.Error(w, "Unable to write to file", http.StatusInternalServerError)
		return
	}

	jsonOutput, _ := json.Marshal(map[string]string{
		"message": "Sensitive value '" + html.EscapeString(string(input)) + "' hashed and stored",
	})
	w.Write(jsonOutput)
}

func doSomething(param string) string {
	// This function simulates some processing of the input parameter
	return param
}
