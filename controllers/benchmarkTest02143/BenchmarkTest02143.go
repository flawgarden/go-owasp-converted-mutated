package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type BenchmarkTest02143 struct{}

func (b *BenchmarkTest02143) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest02143")
	if param == "" {
		param = ""
	}

	bar := doSomething(r, param)

	cookie := http.Cookie{
		Name:     "SomeCookie",
		Value:    bar,
		Secure:   true,
		HttpOnly: true,
		Path:     r.URL.Path,
	}

	http.SetCookie(w, &cookie)
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	responseMessage := fmt.Sprintf("Created cookie: 'SomeCookie': with value: '%s' and secure flag set to: true", htmlEscape(bar))
	w.Write([]byte(responseMessage))
}

func doSomething(r *http.Request, param string) string {
	bar := "safe!"
	map44285 := make(map[string]interface{})
	map44285["keyA-44285"] = "a_Value"
	map44285["keyB-44285"] = param
	map44285["keyC"] = "another_Value"
	bar = map44285["keyB-44285"].(string)
	bar = map44285["keyA-44285"].(string)

	return bar
}

func htmlEscape(str string) string {
	return jsonEscape(str)
}

func jsonEscape(str string) string {
	b, _ := json.Marshal(str)
	return string(b[1 : len(b)-1])
}

func main() {
	http.Handle("/securecookie-00/BenchmarkTest02143", &BenchmarkTest02143{})
	http.ListenAndServe(":8080", nil)
}
