package controllers

import (
	"fmt"
	"net/http"
)

type BenchmarkTest02317Controller struct {
	http.Handler
}

func (c *BenchmarkTest02317Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		c.doPost(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (c *BenchmarkTest02317Controller) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	flag := true
	for name, values := range r.URL.Query() {
		if flag {
			for _, value := range values {
				if value == "BenchmarkTest02317" {
					param = name
					flag = false
					break
				}
			}
		}
	}

	bar := doSomething(r, param)

	w.Header().Set("X-XSS-Protection", "0")
	_, _ = w.Write([]byte(fmt.Sprintf(bar, "a", "b")))
}

func doSomething(r *http.Request, param string) string {
	bar := ""
	num := 196

	if (500/42)+num > 200 {
		bar = param
	} else {
		bar = "This should never happen"
	}

	return bar
}
