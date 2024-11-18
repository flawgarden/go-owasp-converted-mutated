package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

var commonHeaders = []string{"Content-Type", "User-Agent"}

type BenchmarkTest00221 struct{}

func (b *BenchmarkTest00221) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	var param string
	for _, name := range r.Header.Values("X-Custom-Header") {
		if contains(commonHeaders, name) {
			continue
		}
		param = name
		break
	}

	var bar string
	num := 106
	if (7*18)+num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}

	fileName := "/path/to/test/files/" + bar
	var fos *os.File

	defer func() {
		if fos != nil {
			fos.Close()
		}
	}()

	fos, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Couldn't open FileOutputStream on file:", fileName)
		return
	}

	responseMsg := "Now ready to write to file: " + htmlEscape(fileName)
	w.Write([]byte(responseMsg))
}

func contains(slice []string, item string) bool {
	for _, a := range slice {
		if a == item {
			return true
		}
	}
	return false
}

func htmlEscape(s string) string {
	return strconv.Quote(s)[1 : len(strconv.Quote(s))-1]
}

func main() {
	http.Handle("/pathtraver-00/BenchmarkTest00221", &BenchmarkTest00221{})
	http.ListenAndServe(":8080", nil)
}
