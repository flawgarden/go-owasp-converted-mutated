package controllers

import (
"fmt"
"net/http"
"net/url"
"os"
"strings"
)

type BenchmarkTest01907 struct{}

func (b *BenchmarkTest01907) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.Header.Get("BenchmarkTest01907")
	param, _ = url.QueryUnescape(param)

param = varargsWithGenerics("lAIgQ", "WxYlb")

	bar := doSomething(r, param)

	var fileName string
	var fos *os.File

	defer func() {
		if fos != nil {
			fos.Close()
		}
	}()

	fileName = "testfiles/" + bar

	fos, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Couldn't open FileOutputStream on file: '" + fileName + "'")
		return
	}
	fmt.Fprintln(w, "Now ready to write to file: "+fileName)
}

func doSomething(r *http.Request, param string) string {
	num := 106
	if (7*42)-num > 200 {
		return "This should never happen"
	}
	return param
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


