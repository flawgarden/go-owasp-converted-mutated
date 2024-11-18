package controllers

import (
	"net/http"
	"net/url"
)

type BenchmarkTest01080 struct{}

func (b *BenchmarkTest01080) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")

	param := r.Header.Get("BenchmarkTest01080")
	param, _ = url.QueryUnescape(param)

	bar := b.doSomething(r, param)

	r.Context().Value("session").(map[string]interface{})[bar] = "10340"

	w.Write([]byte("Item: '" + sanitizeHTML(bar) + "' with value: 10340 saved in session."))
}

func (b *BenchmarkTest01080) doSomething(r *http.Request, param string) string {
	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[1]         // get the last 'safe' value
	}

	return bar
}

func sanitizeHTML(input string) string {
	// Implement HTML sanitization logic
	return input
}
