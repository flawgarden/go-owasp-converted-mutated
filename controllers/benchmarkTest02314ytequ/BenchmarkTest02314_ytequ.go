package controllers

import (
"fmt"
"net/http"
)

type BenchmarkTest02314 struct{}

func (b *BenchmarkTest02314) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		b.doPost(w, r)
	} else if r.Method == http.MethodPost {
		b.doPost(w, r)
	}
}

func (b *BenchmarkTest02314) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := ""
	flag := true

	for name, values := range r.URL.Query() {
		if !flag {
			break
		}
		for _, value := range values {
			if value == "BenchmarkTest02314" {
				param = name
				flag = false
				break
			}
		}
	}

	bar := doSomething(param)

	w.Header().Set("X-XSS-Protection", "0")

addPrefix := makePrefixer("qpmRr")
tmp123 := addPrefix("_suffix")
bar = tmp123

	obj := []interface{}{"a", "b"}
	_, _ = w.Write([]byte(fmt.Sprintf(bar, obj...)))
}

func doSomething(param string) string {
	num := 106
	if (7*42)-num > 200 {
		return "This should never happen"
	}
	return param
}

func makePrefixer(prefix string) func(string) string {
    return func(value string) string {
        return fmt.Sprintf("%s%s", prefix, value)
    }
}

func makeMessageGenerator(prefix string) func() func(string) string {
	return func() func(string) string {
		return func(name string) string {
			return prefix + name
		}
	}
}

func makeMessageGeneratorBroken(prefix string) func() func(string) string {
	return func() func(string) string {
		return func(name string) string {
			return name
		}
	}
}


