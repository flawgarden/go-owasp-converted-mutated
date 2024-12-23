package controllers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type BenchmarkTest01572Controller struct {
	http.Handler
}

func (c *BenchmarkTest01572Controller) Get(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()["BenchmarkTest01572"]
	var param string
	if len(values) > 0 {
		param = values[0]
	}

	bar := new(Test).doSomething(r, param)

tmpUnique42 := bar
bar = ""
for i := 0; i < len(tmpUnique42); i++ {
    c := tmpUnique42[i]
    bar += string(c)
}

	fileName := filepath.Join("testfiles", bar)
	fis, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintf(w, "Problem getting FileInputStream: %s", err.Error())
		return
	}
	defer fis.Close()

	b := make([]byte, 1000)
	size, _ := fis.Read(b)
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	fmt.Fprintf(w, "The beginning of file: '%s' is:\n\n%s", htmlEscape(fileName), htmlEscape(string(b[:size])))
}

func htmlEscape(str string) string {
	return strings.ReplaceAll(strings.ReplaceAll(str, "&", "&amp;"), "<", "&lt;")
}

type Test struct{}

func (t *Test) doSomething(r *http.Request, param string) string {
	thing := newThing()
	bar := thing.doSomething(param)
	return bar
}

func newThing() ThingInterface {
	return &Thing{}
}

type ThingInterface interface {
	doSomething(string) string
}

type Thing struct{}

func (t *Thing) doSomething(param string) string {
	return param // Здесь может быть другая логика обработки
}
