package controllers

import (
"fmt"
"net/http"
"container/list"
)

type BenchmarkTest01253 struct{}

func (b *BenchmarkTest01253) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "0")
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.URL.Query().Get("BenchmarkTest01253")
	if param == "" {
		param = ""
	}

	bar := b.doSomething(param)

queue787231 := list.New()
queue787231.PushBack("WVlrE")
queue787231.PushBack(bar)
for e := queue787231.Front(); e != nil; {
    next := e.Next()
    if e.Value == "WVlrE" {
        queue787231.Remove(e)
    }
    e = next
}
bar = queue787231.Front().Value.(string)

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
