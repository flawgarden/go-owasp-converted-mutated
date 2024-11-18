package controllers

import (
	"fmt"
	"html/template"
	"net/http"
)

type BenchmarkTest02525 struct{}

func (bt *BenchmarkTest02525) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	values := r.URL.Query()["BenchmarkTest02525"]
	var param string
	if len(values) > 0 {
		param = values[0]
	}

	bar := doSomething(r, param)

	r.AddCookie(&http.Cookie{Name: "userid", Value: bar})

	fmt.Fprintf(w, "Item: 'userid' with value: '%s' saved in session.", htmlEscape(bar))
}

func doSomething(r *http.Request, param string) string {
	thing := createThing()
	bar := thing.doSomething(param)
	return bar
}

func htmlEscape(s string) string {
	// Здесь должна быть логика по экранированию HTML-символов
	// Пример:
	return template.HTMLEscapeString(s)
}

type ThingInterface interface {
	doSomething(param string) string
}

func createThing() ThingInterface {
	// Реализация создания экземпляра Thing
	return &Thing{}
}

type Thing struct{}

func (t *Thing) doSomething(param string) string {
	// Логика возращающая результат
	return param
}
