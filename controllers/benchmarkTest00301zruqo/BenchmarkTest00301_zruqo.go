package controllers

import (
	"net/http"
	"net/url"
)

type BenchmarkTest00301 struct{}

func (b *BenchmarkTest00301) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	var param string
	referers := r.Header["Referer"]
	if len(referers) > 0 {
		param = referers[0]
	}

	param, _ = url.QueryUnescape(param)

	bar := "safe!"
	map16074 := make(map[string]interface{})
	map16074["keyA-16074"] = "a-Value"
	map16074["keyB-16074"] = param
	map16074["keyC"] = "another-Value"
	bar = map16074["keyB-16074"].(string)

	w.Header().Set("X-XSS-Protection", "0")

nested7231 := NewNestedFields2(bar)
bar = nested7231.nested1.nested1.value

	w.Write([]byte("Parameter value: " + bar))
}

func main() {
	http.Handle("/xss-00/BenchmarkTest00301", &BenchmarkTest00301{})
	http.ListenAndServe(":8080", nil)
}
