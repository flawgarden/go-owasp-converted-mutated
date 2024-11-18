package controllers

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"html"
	"net/http"
	"os"
)

func BenchmarkTest00372(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.URL.Query().Get("BenchmarkTest00372")
	if param == "" {
		param = ""
	}

	bar := ""
	num := 106
	if (7*42)-num > 200 {
		bar = "This should never happen"
	} else {
		bar = param
	}

	hashValue := md5.Sum([]byte(bar))
	file, err := os.OpenFile("passwordFile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(w, "Unable to open file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	if _, err := file.WriteString("hash_value=" + base64.StdEncoding.EncodeToString(hashValue[:]) + "\n"); err != nil {
		http.Error(w, "Unable to write to file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Sensitive value '%s' hashed and stored<br/>", html.EscapeString(bar))
	fmt.Fprint(w, "Hash Test executed")
}

func main() {
	http.HandleFunc("/hash-00/BenchmarkTest00372", BenchmarkTest00372)
	http.ListenAndServe(":8080", nil)
}
