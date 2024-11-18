package controllers

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"html"
	"net/http"
	"os"
)

func BenchmarkTest00637(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest00637")
	if param == "" {
		param = ""
	}

	var bar string

	num := 86
	if (7*42)-num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}

	hash := md5.New()
	input := []byte(bar)
	hash.Write(input)
	result := hash.Sum(nil)

	fileTarget, err := os.OpenFile("passwordFile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(w, "Unable to open file", http.StatusInternalServerError)
		return
	}
	defer fileTarget.Close()

	if _, err := fileTarget.WriteString(fmt.Sprintf("hash_value=%s\n", base64.StdEncoding.EncodeToString(result))); err != nil {
		http.Error(w, "Unable to write to file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Sensitive value '%s' hashed and stored<br/>", html.EscapeString(string(input)))
	fmt.Fprintf(w, "Hash Test executed")
}
