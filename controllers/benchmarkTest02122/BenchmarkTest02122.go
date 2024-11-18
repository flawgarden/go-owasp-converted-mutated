package controllers

import (
	"fmt"
	"net/http"
	"text/template"
)

type BenchmarkTest02122 struct{}

func (b *BenchmarkTest02122) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := w.Header()
	response.Set("Content-Type", "text/html;charset=UTF-8")
	param := r.URL.Query().Get("BenchmarkTest02122")
	if param == "" {
		param = ""
	}
	bar := b.doSomething(param)

	response.Set("X-XSS-Protection", "0")
	obj := []string{"a", bar}
	tmpl := `<html><body><p>Formatted like: {{.0}} and {{.1}}.</p></body></html>`
	t := template.Must(template.New("response").Parse(tmpl))
	t.Execute(w, obj)
}

func (b *BenchmarkTest02122) doSomething(param string) string {
	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[1]         // get the last 'safe' value
	}
	return bar
}

func main() {
	http.Handle("/xss-03/BenchmarkTest02122", &BenchmarkTest02122{})
	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
