package controllers

import (
	"fmt"
	"net/http"
)

type BenchmarkTest00546 struct{}

func (b *BenchmarkTest00546) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	responseType := "text/html;charset=UTF-8"
	w.Header().Set("Content-Type", responseType)

	var param string
	flag := true
	for name, values := range r.Form {
		if flag {
			for _, value := range values {
				if value == "BenchmarkTest00546" {
					param = name
					flag = false
					break
				}
			}
		}
	}

	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[1]         // get the last 'safe' value
	}

	w.Header().Set("X-XSS-Protection", "0")
	w.Write([]byte(bar))
}

func main() {
	http.Handle("/xss-01/BenchmarkTest00546", &BenchmarkTest00546{})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server failed:", err)
	}
}
