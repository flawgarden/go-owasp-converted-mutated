package controllers

import (
	"net/http"
	"strings"
)

type BenchmarkTest01268 struct {
}

func (b *BenchmarkTest01268) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := r.URL.Query().Get("BenchmarkTest01268")
	if param == "" {
		param = ""
	}
	bar := b.doSomething(param)
	w.Header().Set("X-XSS-Protection", "0")
	if bar != "" {
		w.Write([]byte(bar))
	}
}

func (b *BenchmarkTest01268) doSomething(param string) string {
	a92384 := param
	b92384 := strings.Builder{}
	b92384.WriteString(a92384)
	b92384.WriteString(" SafeStuff")
	b92384String := b92384.String()
	b92384String = strings.Replace(b92384String, b92384String[len(b92384String)-len("Chars"):], "Chars", 1)
	map92384 := make(map[string]interface{})
	map92384["key92384"] = b92384String
	c92384 := map92384["key92384"].(string)
	d92384 := c92384[:len(c92384)-1]
	e92384 := string([]byte(d92384))
	f92384 := strings.Split(e92384, " ")[0]
	thing := CreateThing()
	bar := thing.DoSomething(f92384)

	return bar
}

// Dummy function that simulates the ThingInterface
type Thing interface {
	DoSomething(input string) string
}

func CreateThing() Thing {
	return &thing{}
}

type thing struct{}

func (t *thing) DoSomething(input string) string {
	return input
}
