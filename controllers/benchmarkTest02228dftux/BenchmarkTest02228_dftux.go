package controllers

import (
	"net/http"
	"strings"
)

type BenchmarkTest02228 struct{}

func (b *BenchmarkTest02228) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.URL.Query().Get("BenchmarkTest02228")

var ptr123 *string = new(string)
var ptr234 *string = new(string)
*ptr123 = param
ptr123 = ptr234
param = *ptr123

	bar := doSomething(param)

	w.Header().Set("X-XSS-Protection", "0")
	w.Write([]byte(bar))
}

func doSomething(param string) string {
	var sb strings.Builder
	sb.WriteString(param)
	bar := sb.String() + "_SafeStuff"
	return bar
}

func main() {
	http.Handle("/xss-04/BenchmarkTest02228", &BenchmarkTest02228{})
	http.ListenAndServe(":8080", nil)
}

func addSuffix(s *string, suf string) {
	*s = *s + suf
}

func addSuffixDoublePointer(s **string, suf *string) {
	**s = **s + *suf
}

func addSuffixDoublePointerBroken(s **string, suf *string) {
	*s = new(string)
	**s = **s + *suf
}

func swapStrings(a, b *string) {
	temp := *a
	*a = *b
	*b = temp
}

func removeSpaces(s *string) {
    *s = strings.ReplaceAll(*s, " ", "")
}


