package controllers

import (
	"net/http"
)

type BenchmarkTest01599 struct{}

func (b *BenchmarkTest01599) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "0")
	values := r.URL.Query()["BenchmarkTest01599"]
	var param string
	if len(values) > 0 {
		param = values[0]
	} else {
		param = ""
	}

	bar := b.doSomething(r, param)
	w.Write([]byte(bar))
}

func (b *BenchmarkTest01599) doSomething(r *http.Request, param string) string {
	bar := "safe!"
	map20521 := make(map[string]interface{})
	map20521["keyA-20521"] = "a_Value"
	map20521["keyB-20521"] = param
	map20521["keyC"] = "another_Value"
	bar = map20521["keyB-20521"].(string)
	bar = map20521["keyA-20521"].(string)

	return bar
}

func main() {
	http.Handle("/xss-03/BenchmarkTest01599", &BenchmarkTest01599{})
	http.ListenAndServe(":8080", nil)
}
