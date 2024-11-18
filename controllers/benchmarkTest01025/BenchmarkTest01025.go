package controllers

import (
	"net/http"
	"net/url"
	"os"
	"strings"
)

type BenchmarkTest01025 struct{}

func (b *BenchmarkTest01025) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	if r.Header.Get("BenchmarkTest01025") != "" {
		param = r.Header.Get("BenchmarkTest01025")
	}

	param, _ = url.QueryUnescape(param)

	bar := b.doSomething(r, param)

	fileTarget := os.Getenv("TESTFILES_DIR") + "/" + bar
	w.Write([]byte("Access to file: '" + sanitize(fileTarget) + "' created."))
	if _, err := os.Stat(fileTarget); os.IsNotExist(err) {
		w.Write([]byte(" But file doesn't exist yet."))
	} else {
		w.Write([]byte(" And file already exists."))
	}
}

func (b *BenchmarkTest01025) doSomething(r *http.Request, param string) string {
	thing := createThing()
	bar := thing.doSomething(param)
	return bar
}

type ThingInterface interface {
	doSomething(param string) string
}

func createThing() ThingInterface {
	return &Thing{}
}

type Thing struct{}

func (t *Thing) doSomething(param string) string {
	return strings.TrimSpace(param)
}

func sanitize(input string) string {
	return strings.ReplaceAll(input, "<", "&lt;")
}
