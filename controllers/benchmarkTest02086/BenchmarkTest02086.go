package controllers

import (
	"net/http"
	"net/url"
	"strings"
)

type BenchmarkTest02086 struct{}

func (b *BenchmarkTest02086) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	if headers := r.Header["BenchmarkTest02086"]; len(headers) > 0 {
		param = headers[0]
	}

	param, _ = url.QueryUnescape(param)

	bar := doSomething(param)

	r.Header.Set("userid", bar)

	w.Write([]byte("Item: 'userid' with value: '" + escapeForHTML(bar) + "' saved in session."))
}

func doSomething(param string) string {
	var bar string
	guess := "ABC"
	switchTarget := guess[1] // condition 'B', which is safe

	switch switchTarget {
	case 'A':
		bar = param
	case 'B':
		bar = "bob"
	case 'C', 'D':
		bar = param
	default:
		bar = "bob's your uncle"
	}

	return bar
}

func escapeForHTML(s string) string {
	m := map[string]string{
		"&": "&amp;",
		"<": "&lt;",
		">": "&gt;",
		"'": "&#39;",
		`"`: `&quot;`,
	}
	for old, new := range m {
		s = strings.ReplaceAll(s, old, new)
	}
	return s
}
