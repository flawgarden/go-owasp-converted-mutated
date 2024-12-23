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

bar = combineStrings("CIBWP", "GpnbI")

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

func getFirstString(lines ...string) string {
    return getStringWithIndex(0, lines...)
}

func getStringWithIndex(ind int, lines ...string) string {
    return lines[ind]
}

func getFirstStringFromArray(lines ...string) string {
    return lines[0]
}

func varargsWithGenerics[T any](elements ...T) T {
    return elements[0]
}

func combineStrings(strs ...string) string {
    return strings.Join(strs, ", ")
}


