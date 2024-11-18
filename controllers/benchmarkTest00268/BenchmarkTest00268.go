package controllers

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

type BenchmarkTest00268Controller struct {
	http.ResponseWriter
	*http.Request
}

func (c *BenchmarkTest00268Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	var param string
	if headers := r.Header["BenchmarkTest00268"]; len(headers) > 0 {
		param = headers[0]
	}
	param, _ = url.QueryUnescape(param)

	bar := ""
	num := 106
	if (7*42)-num > 200 {
		bar = "This should never happen"
	} else {
		bar = param
	}

	hash := sha1.New()
	input := []byte(bar)
	hash.Write(input)
	result := hash.Sum(nil)

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

	jsonResponse, err := json.Marshal(map[string]string{
		"message": fmt.Sprintf("Sensitive value '%s' hashed and stored", htmlEscape(string(input))),
	})
	if err != nil {
		http.Error(w, "Unable to create JSON response", http.StatusInternalServerError)
		return
	}

	w.Write(jsonResponse)
}

func htmlEscape(input string) string {
	// Implement a simple HTML escape if needed
	return input
}
