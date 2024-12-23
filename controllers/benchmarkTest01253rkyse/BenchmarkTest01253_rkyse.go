package controllers

import (
"fmt"
"net/http"
"sync"
)

type BenchmarkTest01253 struct{}

func (b *BenchmarkTest01253) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "0")
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.URL.Query().Get("BenchmarkTest01253")
	if param == "" {
		param = ""
	}

message123 := make(chan string, 1)
message123 <- param

var wg sync.WaitGroup
wg.Add(1)
go func() {
    defer wg.Done()
    rmsg := <-message123
    message123 <- rmsg + "constant_string"
}()

wg.Wait()

param = <-message123

	bar := b.doSomething(param)

	obj := []interface{}{"a", "b"}
	_, err := w.Write([]byte(fmt.Sprintf(bar, obj...)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (b *BenchmarkTest01253) doSomething(param string) string {
	bar := "safe!"
	map15481 := make(map[string]interface{})
	map15481["keyA-15481"] = "a-Value"
	map15481["keyB-15481"] = param
	map15481["keyC"] = "another-Value"
	bar = map15481["keyB-15481"].(string)

	return bar
}
