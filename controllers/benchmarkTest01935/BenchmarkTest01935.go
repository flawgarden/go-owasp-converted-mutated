package controllers

import (
	"fmt"
	"net/http"
	"net/url"
)

type BenchmarkTest01935 struct{}

func (b *BenchmarkTest01935) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		b.doPost(w, r)
	} else if r.Method == http.MethodPost {
		b.doPost(w, r)
	}
}

func (b *BenchmarkTest01935) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	if r.Header.Get("BenchmarkTest01935") != "" {
		param = r.Header.Get("BenchmarkTest01935")
	}

	param, _ = url.QueryUnescape(param)

	bar := doSomething(param)

	cookie := http.Cookie{
		Name:     "SomeCookie",
		Value:    bar,
		Secure:   true,
		HttpOnly: true,
		Path:     r.RequestURI,
	}
	http.SetCookie(w, &cookie)

	fmt.Fprintf(w, "Created cookie: 'SomeCookie': with value: '%s' and secure flag set to: true", htmlEscape(bar))
}

func doSomething(param string) string {
	a17785 := param
	b17785 := a17785 + " SafeStuff"
	b17785 = b17785[:len(b17785)-len("Chars")] + "Chars"
	map17785 := map[string]interface{}{"key17785": b17785}
	c17785 := map17785["key17785"].(string)
	d17785 := c17785[:len(c17785)-1]
	e17785 := string(Base64Decode(Base64Encode([]byte(d17785))))
	f17785 := splitOnSpace(e17785)[0]

	thing := createThing()
	bar := thing.doSomething(f17785)

	return bar
}

func htmlEscape(input string) string {
	// implement HTML escape
	return input
}

func Base64Encode(data []byte) []byte {
	// implement Base64 encode
	return data
}

func Base64Decode(data []byte) []byte {
	// implement Base64 decode
	return data
}

func splitOnSpace(input string) []string {
	// implement split on space
	return []string{input}
}

func createThing() ThingInterface {
	// implement create thing
	return nil
}

type ThingInterface interface {
	doSomething(input string) string
}
