package controllers

import (
	"fmt"
	"net/http"
)

type BenchmarkTest02239 struct{}

func (t *BenchmarkTest02239) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "0")
	r.ParseForm()

	param := r.FormValue("BenchmarkTest02239")
	bar := doSomething(param)

	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	w.Write([]byte(bar))
}

func doSomething(param string) string {
	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the last 'safe' value
	}
	return bar
}

func main() {
	http.Handle("/xss-04/BenchmarkTest02239", &BenchmarkTest02239{})
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
