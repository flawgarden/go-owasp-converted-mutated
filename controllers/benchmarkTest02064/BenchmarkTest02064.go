package controllers

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type BenchmarkTest02064 struct{}

func (b *BenchmarkTest02064) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	if headers := r.Header["BenchmarkTest02064"]; len(headers) > 0 {
		param = headers[0]
	}

	param, _ = url.QueryUnescape(param)

	bar := doSomething(param)

	cookie := http.Cookie{
		Name:     "SomeCookie",
		Value:    bar,
		Path:     r.RequestURI,
		Secure:   true,
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie)

	response := map[string]string{
		"message": "Created cookie: 'SomeCookie': with value: '" + bar + "' and secure flag set to: true",
	}
	json.NewEncoder(w).Encode(response)
}

func doSomething(param string) string {
	bar := "safe!"
	map96496 := make(map[string]interface{})
	map96496["keyA-96496"] = "a_Value"
	map96496["keyB-96496"] = param
	map96496["keyC"] = "another_Value"
	bar = map96496["keyB-96496"].(string)
	bar = map96496["keyA-96496"].(string)

	return bar
}

func main() {
	http.Handle("/securecookie-00/BenchmarkTest02064", &BenchmarkTest02064{})
	http.ListenAndServe(":8080", nil)
}
