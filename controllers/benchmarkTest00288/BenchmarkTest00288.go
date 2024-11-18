package controllers

import (
	"fmt"
	"net/http"
	"strings"
)

type BenchmarkTest00288 struct{}

func (b *BenchmarkTest00288) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := ""

	if referer := r.Header.Get("Referer"); referer != "" {
		param = referer
	}

	param = strings.TrimSpace(param)

	bar := "safe!"
	map34285 := make(map[string]interface{})
	map34285["keyA-34285"] = "a_Value"
	map34285["keyB-34285"] = param
	map34285["keyC"] = "another_Value"

	bar = map34285["keyB-34285"].(string)
	bar = map34285["keyA-34285"].(string)

	w.Header().Set("X-XSS-Protection", "0")
	length := 1
	if bar != "" {
		length = len(bar)
		w.Write([]byte(bar)[:length])
	}
}

func main() {
	http.Handle("/xss-00/BenchmarkTest00288", &BenchmarkTest00288{})
	fmt.Println("Server is listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
