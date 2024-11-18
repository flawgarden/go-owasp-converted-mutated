package controllers

import (
	"net/http"
)

type BenchmarkTest02712 struct{}

func (b *BenchmarkTest02712) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		b.doPost(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (b *BenchmarkTest02712) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	param := r.URL.Query().Get("BenchmarkTest02712")
	bar := b.doSomething(r, param)

	w.Header().Set("X-XSS-Protection", "0")
	_, _ = w.Write([]byte("Parameter value: " + bar))
}

func (b *BenchmarkTest02712) doSomething(r *http.Request, param string) string {
	bar := "safe!"
	map52216 := make(map[string]interface{})
	map52216["keyA-52216"] = "a_Value"
	map52216["keyB-52216"] = param
	map52216["keyC"] = "another_Value"
	bar = map52216["keyB-52216"].(string)
	bar = map52216["keyA-52216"].(string)

	return bar
}

func main() {
	http.Handle("/xss-05/BenchmarkTest02712", &BenchmarkTest02712{})
	http.ListenAndServe(":8080", nil)
}
