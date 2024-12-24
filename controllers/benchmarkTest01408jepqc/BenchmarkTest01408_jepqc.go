package controllers

import (
"fmt"
"net/http"
"os"
"strings"
)

type BenchmarkTest01408 struct{}

func (b *BenchmarkTest01408) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	param := ""
	flag := true
	for name, values := range r.URL.Query() {
		if !flag {
			break
		}
		for _, value := range values {
			if value == "BenchmarkTest01408" {
				param = name
				flag = false
			}
		}
	}

param = getFirstStringFromArray(param, param)

	bar := b.doSomething(r, param)

	fileName := fmt.Sprintf("/path/to/test/files/%s", bar)
	var fos *os.File
	var err error

	defer func() {
		if fos != nil {
			fos.Close()
		}
	}()

	fos, err = os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Couldn't open FileOutputStream on file: '" + fileName + "'")
		return
	}

	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	w.Write([]byte("Now ready to write to file: " + fileName))
}

func (b *BenchmarkTest01408) doSomething(r *http.Request, param string) string {
	thing := ThingFactory{}.createThing()
	bar := thing.doSomething(param)
	return bar
}

type ThingInterface interface {
	doSomething(param string) string
}

type ThingFactory struct{}

func (f ThingFactory) createThing() ThingInterface {
	return &Thing{}
}

type Thing struct{}

func (t *Thing) doSomething(param string) string {
	return param // Example transformation, replace with actual logic
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


