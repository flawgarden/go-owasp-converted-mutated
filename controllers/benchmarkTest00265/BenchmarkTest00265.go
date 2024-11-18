package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type BenchmarkTest00265 struct{}

func (b *BenchmarkTest00265) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		b.doPost(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (b *BenchmarkTest00265) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	if headers := r.Header["BenchmarkTest00265"]; len(headers) > 0 {
		param = headers[0]
	}

	param, _ = url.QueryUnescape(param)

	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the last 'safe' value
	}

	fileName := fmt.Sprintf("./testfiles/%s", bar)
	var fos *os.File

	defer func() {
		if fos != nil {
			fos.Close()
		}
	}()

	var err error
	fos, err = os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Couldn't open FileOutputStream on file: '%s'\n", fileName)
		return
	}

	_, _ = w.Write([]byte(fmt.Sprintf("Now ready to write to file: %s", htmlEscape(fileName))))
}

func htmlEscape(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(s, "&", "&amp;"), "<", "&lt;")
}
