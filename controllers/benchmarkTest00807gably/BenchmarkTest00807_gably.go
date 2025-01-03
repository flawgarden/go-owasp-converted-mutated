package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type BenchmarkTest00807 struct{}

func (bt *BenchmarkTest00807) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	r.ParseForm()
	queryString := r.URL.RawQuery
	paramval := "BenchmarkTest00807="
	paramLoc := strings.Index(queryString, paramval)

	if paramLoc == -1 {
		http.Error(w, fmt.Sprintf("getQueryString() couldn't find expected parameter '%s' in query string.", "BenchmarkTest00807"), http.StatusBadRequest)
		return
	}

	param := queryString[paramLoc+len(paramval):]

var ptr123 *string = new(string)
var ptr234 *string = new(string)
*ptr234 = param
ptr123 = ptr234
param = *ptr123

	ampersandLoc := strings.Index(param, "&")
	if ampersandLoc != -1 {
		param = param[:ampersandLoc]
	}
	param, _ = url.QueryUnescape(param)

	bar := ""
	if param != "" {
		bar = strings.Split(param, " ")[0]
	}

	w.Header().Set("X-XSS-Protection", "0")
	output := fmt.Sprintf(bar)
	w.Write([]byte(output))
}

func main() {
	http.Handle("/xss-01/BenchmarkTest00807", &BenchmarkTest00807{})
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


