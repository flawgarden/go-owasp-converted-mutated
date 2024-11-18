package controllers

import (
	"fmt"
	"net/http"
)

type BenchmarkTest00376 struct{}

func (b *BenchmarkTest00376) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.URL.Query().Get("BenchmarkTest00376")
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

	w.Header().Set("X-XSS-Protection", "0")
	obj := []interface{}{"a", bar}
	output := fmt.Sprintf("Formatted like: %s and %s.", obj[0], obj[1])
	w.Write([]byte(output))
}

func main() {
	http.Handle("/xss-00/BenchmarkTest00376", &BenchmarkTest00376{})
	http.ListenAndServe(":8080", nil)
}
