package controllers

import (
	"net/http"
	"text/template"
)

type BenchmarkTest01251Controller struct {
	http.Handler
}

func (c *BenchmarkTest01251Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest01251")
	if param == "" {
		param = ""
	}

	bar := new(Test).doSomething(param)

	w.Header().Set("X-XSS-Protection", "0")
	obj := []string{"a", bar}
	tmpl := `<html>
<body>
<p>
Formatted like: {{ .[0] }} and {{ .[1] }}.
</p>
</body>
</html>`
	t := template.Must(template.New("webpage").Parse(tmpl))
	t.Execute(w, obj)
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	bar := param // Здесь может быть использована библиотека для экранирования
	return bar
}
