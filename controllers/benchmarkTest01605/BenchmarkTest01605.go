package controllers

import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01605 struct{}

func (b *BenchmarkTest01605) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest01605")

	bar := new(Test).doSomething(r, param)

	cookie := &http.Cookie{
		Name:     "SomeCookie",
		Value:    bar,
		HttpOnly: true,
		Secure:   true,
		Path:     r.URL.Path,
	}
	http.SetCookie(w, cookie)

	response := fmt.Sprintf("Created cookie: 'SomeCookie': with value: '%s' and secure flag set to: true", bar)
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	w.Write([]byte(response))
}

type Test struct{}

func (t *Test) doSomething(r *http.Request, param string) string {
	bar := "safe!"
	mapTest := map[string]interface{}{
		"keyA-97880": "a_Value",
		"keyB-97880": param,
		"keyC":       "another_Value",
	}
	bar = mapTest["keyB-97880"].(string)
	bar = mapTest["keyA-97880"].(string)

	return bar
}
