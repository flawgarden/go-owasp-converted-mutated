package controllers

import (
	"encoding/base64"
	"fmt"
	"net/http"
)

type BenchmarkTest00473 struct{}

func (b *BenchmarkTest00473) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "0")
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.URL.Query().Get("BenchmarkTest00473")
	bar := ""
	if param != "" {
		decoded, _ := base64.StdEncoding.DecodeString(base64.StdEncoding.EncodeToString([]byte(param)))
		bar = string(decoded)
	}

	obj := []interface{}{"a", "b"}
	w.Write([]byte(fmt.Sprintf(bar, obj...)))
}
