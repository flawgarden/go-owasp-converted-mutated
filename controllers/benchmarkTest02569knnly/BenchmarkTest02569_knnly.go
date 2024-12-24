package controllers

import (
	"net/http"
	"os"
	"strings"
)

type BenchmarkTest02569Controller struct {
	http.ResponseWriter
	*http.Request
}

func (c *BenchmarkTest02569Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.ResponseWriter = w
	c.Request = r
	c.doPost()
}

func (c *BenchmarkTest02569Controller) doPost() {
	c.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	queryString := c.Request.URL.RawQuery

str123 := "const_string"
swapStrings(&queryString, &str123)
queryString = str123

	paramVal := "BenchmarkTest02569="
	paramLoc := strings.Index(queryString, paramVal)

	if paramLoc == -1 {
		c.ResponseWriter.Write([]byte("getQueryString() couldn't find expected parameter 'BenchmarkTest02569' in query string."))
		return
	}

	param := queryString[paramLoc+len(paramVal):]
	ampersandLoc := strings.Index(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramVal) : paramLoc+ampersandLoc]
	}

	bar := doSomething(param)
	fileName := bar

	fos, err := os.Create(fileName)
	if err != nil {
		c.ResponseWriter.Write([]byte("Couldn't open FileOutputStream on file: '" + fileName + "'"))
		return
	}
	defer fos.Close()

	c.ResponseWriter.Write([]byte("Now ready to write to file: " + fileName))
}

func doSomething(param string) string {
	bar := ""
	switchTarget := 'C'

	switch switchTarget {
	case 'A':
		bar = param
	case 'B':
		bar = "bobs_your_uncle"
	case 'C', 'D':
		bar = param
	default:
		bar = "bobs_your_uncle"
	}

	return bar
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


