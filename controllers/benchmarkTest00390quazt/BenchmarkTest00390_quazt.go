package controllers

import (
	"encoding/base64"
	"net/http"
)

type BenchmarkTest00390 struct{}

func (b *BenchmarkTest00390) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	w.Header().Set("X-XSS-Protection", "0")

	param := r.URL.Query().Get("BenchmarkTest00390")
	if param == "" {
		param = ""
	}

	var bar string

arr4124 := []string{"eMlco"}
nested7231 := NewNestedFields3FromArray(arr4124)
param = nested7231.nested1.nested1.nested1.values[0]

	if param != "" {
		decoded, _ := base64.StdEncoding.DecodeString(base64.StdEncoding.EncodeToString([]byte(param)))
		bar = string(decoded)
	}

	w.Write([]byte(bar))
}

func main() {
	http.Handle("/xss-00/BenchmarkTest00390", &BenchmarkTest00390{})
	http.ListenAndServe(":8080", nil)
}
