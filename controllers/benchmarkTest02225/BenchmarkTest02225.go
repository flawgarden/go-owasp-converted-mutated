package controllers

import (
	"fmt"
	"net/http"
)

type BenchmarkTest02225Controller struct{}

func (c *BenchmarkTest02225Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "0")
	r.ParseForm()
	param := r.FormValue("BenchmarkTest02225")
	bar := doSomething(r, param)

	obj := []interface{}{"a", "b"}
	fmt.Fprintf(w, bar, obj...)
}

func doSomething(r *http.Request, param string) string {
	bar := "safe!"
	map73885 := make(map[string]interface{})
	map73885["keyA-73885"] = "a-Value"
	map73885["keyB-73885"] = param
	map73885["keyC"] = "another-Value"
	bar = map73885["keyB-73885"].(string)

	return bar
}
