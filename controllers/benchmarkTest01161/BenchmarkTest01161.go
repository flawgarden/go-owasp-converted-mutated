package controllers

import (
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

type BenchmarkTest01161Controller struct {
	http.Handler
}

func (c *BenchmarkTest01161Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		c.doPost(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (c *BenchmarkTest01161Controller) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.Header.Get("BenchmarkTest01161")
	param = decode(param)

	bar := c.doSomething(param)

	fileName := ""
	var fos *os.File

	defer func() {
		if fos != nil {
			fos.Close()
		}
	}()

	fileName = filepath.Join("testfiles", bar)

	fos, _ = os.Create(fileName)
	if fos != nil {
		_, _ = w.Write([]byte("Now ready to write to file: " + htmlEscape(fileName)))
	}
}

func (c *BenchmarkTest01161Controller) doSomething(param string) string {
	bar := ""
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the param value
	}
	return bar
}

func decode(param string) string {
	decoded, _ := url.QueryUnescape(param)
	return decoded
}

func htmlEscape(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(s, "&", "&amp;"), "<", "&lt;")
}
