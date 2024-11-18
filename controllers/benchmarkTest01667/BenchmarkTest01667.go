package controllers

import (
	"net/http"
)

type BenchmarkTest01667 struct{}

func (b *BenchmarkTest01667) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	queryString := r.URL.Query()
	param := queryString.Get("BenchmarkTest01667")
	if param == "" {
		http.Error(w, "getQueryString() couldn't find expected parameter 'BenchmarkTest01667' in query string.", http.StatusBadRequest)
		return
	}

	bar := b.doSomething(param)
	w.Header().Set("X-XSS-Protection", "0")
	w.Write([]byte(bar))
}

func (b *BenchmarkTest01667) doSomething(param string) string {
	bar := "safe!"
	map96771 := make(map[string]interface{})
	map96771["keyA-96771"] = "a-Value"
	map96771["keyB-96771"] = param
	map96771["keyC"] = "another-Value"
	bar, _ = map96771["keyB-96771"].(string)
	return bar
}

func main() {
	http.Handle("/xss-03/BenchmarkTest01667", &BenchmarkTest01667{})
	http.ListenAndServe(":8080", nil)
}
