package controllers

import (
	"encoding/json"
	"net/http"
	"strings"
)

type BenchmarkTest02324Controller struct {
	http.Handler
}

func (c *BenchmarkTest02324Controller) Get(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := ""
	flag := true

	for name, values := range r.Form {
		if flag {
			for _, value := range values {
				if value == "BenchmarkTest02324" {
					param = name
					flag = false
					break
				}
			}
		}
	}

	bar := doSomething(param)

	w.Header().Set("X-XSS-Protection", "0")
	output := struct {
		Formatted string `json:"formatted"`
	}{
		Formatted: "Formatted like: " + bar + " and b.",
	}
	response, _ := json.Marshal(output)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func doSomething(param string) string {
	bar := param
	if param != "" && len(param) > 1 {
		sbxyz12198 := strings.Builder{}
		sbxyz12198.WriteString(param)
		bar = sbxyz12198.String()[:len(param)-1] + "Z"
	}
	return bar
}
