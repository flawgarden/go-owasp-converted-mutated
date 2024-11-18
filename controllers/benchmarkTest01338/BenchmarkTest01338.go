package controllers

import (
	"net/http"
)

type BenchmarkTest01338 struct{}

func (b *BenchmarkTest01338) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest01338")
	bar := new(Test).doSomething(r, param)

	w.Header().Set("X-XSS-Protection", "0")
	w.Write([]byte(bar))
}

type Test struct{}

func (t *Test) doSomething(r *http.Request, param string) string {
	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the last 'safe' value
	}
	return bar
}

func main() {
	http.Handle("/xss-02/BenchmarkTest01338", &BenchmarkTest01338{})
	http.ListenAndServe(":8080", nil)
}
