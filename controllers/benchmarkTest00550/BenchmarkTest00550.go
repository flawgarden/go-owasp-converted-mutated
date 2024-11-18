package controllers

import (
	"net/http"
)

type BenchmarkTest00550 struct{}

func (bt *BenchmarkTest00550) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := w.Header()
	response.Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	flag := true
	names := r.URL.Query()

	for name, values := range names {
		if flag {
			for _, value := range values {
				if value == "BenchmarkTest00550" {
					param = name
					flag = false
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

	response.Set("X-XSS-Protection", "0")
	w.Write([]byte(bar))
}
