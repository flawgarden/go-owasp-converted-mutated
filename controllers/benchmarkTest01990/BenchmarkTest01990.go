package controllers

import (
	"fmt"
	"net/http"
	"os"
)

type BenchmarkTest01990 struct{}

func (b *BenchmarkTest01990) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	for name := range r.Header {
		if isCommonHeader(name) {
			continue
		}
		param = name
		break
	}

	bar := doSomething(param)

	fileName := fmt.Sprintf("/path/to/testfiles/%s", bar)
	fos, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Couldn't open FileOutputStream on file: '" + fileName + "'")
		return
	}
	defer fos.Close()

	_, err = w.Write([]byte("Now ready to write to file: " + htmlEscape(fileName)))
	if err != nil {
		fmt.Println("Couldn't write to response: " + err.Error())
	}
}

func isCommonHeader(name string) bool {
	commonHeaders := []string{"Content-Type", "User-Agent"}
	for _, header := range commonHeaders {
		if name == header {
			return true
		}
	}
	return false
}

func doSomething(param string) string {
	bar, guess := "", "ABC"
	switch guess[1] {
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

func htmlEscape(s string) string {
	return fmt.Sprintf("%s", s) // здесь можно реализовать экранирование для HTML
}
