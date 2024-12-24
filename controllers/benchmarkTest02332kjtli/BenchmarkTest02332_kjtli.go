package controllers

import (
	"fmt"
	"net/http"
)

type BenchmarkTest02332 struct{}

func (b *BenchmarkTest02332) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	var param string
	flag := true
	for name, values := range r.URL.Query() {
		if flag {
			for _, value := range values {

ii := NewInstanceInitializer(name)
name = ii.list[0]

				if value == "BenchmarkTest02332" {
					param = name
					flag = false
					break
				}
			}
		}
	}

	bar := doSomething(r, param)

	w.Header().Set("X-XSS-Protection", "0")
	w.Write([]byte(bar))
}

func doSomething(r *http.Request, param string) string {
	bar := ""
	num := 106
	if (7*42)-num > 200 {
		bar = "This should never happen"
	} else {
		bar = param
	}
	return bar
}

func main() {
	http.Handle("/xss-04/BenchmarkTest02332", &BenchmarkTest02332{})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
