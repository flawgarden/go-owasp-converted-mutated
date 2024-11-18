package controllers

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"html"
	"net/http"
	"os"
)

type BenchmarkTest00536 struct{}

func (b *BenchmarkTest00536) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	flag := true
	names := r.URL.Query()

	for name, values := range names {
		if flag {
			for _, value := range values {
				if value == "BenchmarkTest00536" {
					param = name
					flag = false
					break
				}
			}
		}
	}

	bar := "safe!"
	map28714 := make(map[string]interface{})
	map28714["keyA-28714"] = "a-Value"
	map28714["keyB-28714"] = param
	map28714["keyC"] = "another-Value"
	bar = map28714["keyB-28714"].(string)

	hash := md5.Sum([]byte(bar))
	result := base64.StdEncoding.EncodeToString(hash[:])

	file, _ := os.OpenFile("passwordFile.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	defer file.Close()
	file.WriteString(fmt.Sprintf("hash_value=%s\n", result))

	fmt.Fprintf(w, "Sensitive value '%s' hashed and stored<br/>", html.EscapeString(bar))
	fmt.Fprint(w, "Hash Test executed")
}

func main() {
	http.Handle("/hash-00/BenchmarkTest00536", &BenchmarkTest00536{})
	http.ListenAndServe(":8080", nil)
}
