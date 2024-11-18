package controllers

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"html"
	"net/http"
)

type BenchmarkTest00462 struct{}

func (b *BenchmarkTest00462) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.Redirect(w, r, r.URL.String(), http.StatusSeeOther)
		return
	}

	r.ParseForm()
	param := r.FormValue("BenchmarkTest00462")

	bar := ""
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the param value
	}

	h := sha1.New()
	h.Write([]byte(bar))
	result := h.Sum(nil)

	encoded := base64.StdEncoding.EncodeToString(result)

	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	fmt.Fprintf(w, "Sensitive value '%s' hashed and stored<br/>", html.EscapeString(bar))
	fmt.Fprintf(w, "hash_value=%s\n", encoded)
	fmt.Fprintln(w, "Hash Test executed")
}
