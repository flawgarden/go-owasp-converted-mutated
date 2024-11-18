package controllers

import (
	"net/http"
	"net/url"
	"text/template"
)

type BenchmarkTest01920 struct{}

func (b *BenchmarkTest01920) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	if referrer := r.Header.Get("Referer"); referrer != "" {
		param = referrer
	}

	decodedParam, _ := url.QueryUnescape(param)
	bar := doSomething(r, decodedParam)

	w.Header().Set("X-XSS-Protection", "0")
	tmpl := template.Must(template.New("output").Parse("Formatted like: {{.Bar}} and {{.B}}."))
	tmpl.Execute(w, map[string]interface{}{
		"Bar": bar,
		"B":   "b",
	})
}

func doSomething(r *http.Request, param string) string {
	bar := ""
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the param value
	}
	return bar
}

func main() {
	http.Handle("/xss-03/BenchmarkTest01920", &BenchmarkTest01920{})
	http.ListenAndServe(":8080", nil)
}
