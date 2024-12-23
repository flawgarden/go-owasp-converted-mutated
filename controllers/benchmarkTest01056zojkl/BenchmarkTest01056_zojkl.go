package controllers

import (
	"net/http"
	"net/url"
)

type BenchmarkTest01056 struct{}

func (b *BenchmarkTest01056) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	if r.Header.Get("Referer") != "" {
		param = r.Header.Get("Referer")

tmpArrayUnique42 := []string{"", "", ""}
tmpArrayUnique42[0] = param
ah := NewArrayHolderWithValues(tmpArrayUnique42)
param = ah.Values[0]

	}

	param, _ = url.QueryUnescape(param)

	bar := b.doSomething(r, param)

	w.Header().Set("X-XSS-Protection", "0")
	length := 1
	if bar != "" {
		length = len(bar)
		w.Write([]byte(bar)[:length])
	}
}

func (b *BenchmarkTest01056) doSomething(r *http.Request, param string) string {
	bar := ""

	num := 106
	if (7*42)-num > 200 {
		bar = "This should never happen"
	} else {
		bar = param
	}

	return bar
}
