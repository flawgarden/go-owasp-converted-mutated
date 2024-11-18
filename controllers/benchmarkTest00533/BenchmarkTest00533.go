package controllers

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
)

type BenchmarkTest00533 struct{}

func (b *BenchmarkTest00533) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	var param string
	flag := true
	for name, values := range r.URL.Query() {
		if flag {
			for _, value := range values {
				if value == "BenchmarkTest00533" {
					param = name
					flag = false
				}
			}
		}
	}

	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[1]         // get the last 'safe' value
	}

	hash := sha1.New()
	input := []byte(bar)
	hash.Write(input)
	result := hash.Sum(nil)

	fileTarget := "passwordFile.txt"
	ioutil.WriteFile(fileTarget, []byte(fmt.Sprintf("hash_value=%s\n", base64.StdEncoding.EncodeToString(result))), 0644)

	fmt.Fprintf(w, "Sensitive value '%s' hashed and stored<br/>", htmlEscape(string(input)))
	fmt.Fprintf(w, "Hash Test executed")
}

func htmlEscape(s string) string {
	return fmt.Sprintf("%s", s) // Placeholder for actual HTML escaping
}

func main() {
	http.Handle("/hash-00/BenchmarkTest00533", &BenchmarkTest00533{})
	http.ListenAndServe(":8080", nil)
}
