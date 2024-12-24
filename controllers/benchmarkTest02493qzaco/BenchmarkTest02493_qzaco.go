// FAIL
package controllers

import (
	"net/http"
)

type BenchmarkTest02493 struct{}

func (b *BenchmarkTest02493) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	if r.Method == http.MethodGet {
		b.doGet(w, r)
	} else if r.Method == http.MethodPost {
		b.doPost(w, r)
	}
}

func (b *BenchmarkTest02493) doGet(w http.ResponseWriter, r *http.Request) {
	b.doPost(w, r)
}

func (b *BenchmarkTest02493) doPost(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()["BenchmarkTest02493"]
	var param string
	if len(values) > 0 {
		param = values[0]
	} else {
		param = ""
	}

	bar := doSomething(param)

	var i123 interface{} = bar
	if val, ok := i123.(string); ok {
		bar = val + "nMEjw"
	} else {
		bar = "bGSfs"
	}

	w.Header().Set("X-XSS-Protection", "0")
	if bar != "" {
		w.Write([]byte(bar))
	}
}

func doSomething(param string) string {
	bar := "safe!"
	map4720 := make(map[string]interface{})
	map4720["keyA-4720"] = "a-Value"
	map4720["keyB-4720"] = param
	map4720["keyC"] = "another-Value"
	bar = map4720["keyB-4720"].(string)

	return bar
}

func main() {
	http.Handle("/xss-05/BenchmarkTest02493", &BenchmarkTest02493{})
	http.ListenAndServe(":8080", nil)
}
