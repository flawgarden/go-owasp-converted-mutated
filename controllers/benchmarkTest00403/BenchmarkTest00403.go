package controllers

import (
	"encoding/json"
	"net/http"
)

type BenchmarkTest00403 struct{}

func (b *BenchmarkTest00403) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.URL.Query().Get("BenchmarkTest00403")
	if param == "" {
		param = ""
	}

	bar := ""
	num := 106
	if (7*18)+num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}

	cookie := &http.Cookie{
		Name:     "SomeCookie",
		Value:    bar,
		HttpOnly: true,
		Path:     r.RequestURI,
		Secure:   false,
	}
	http.SetCookie(w, cookie)

	response := map[string]string{
		"message": "Created cookie: 'SomeCookie': with value: '" + bar + "' and secure flag set to: false",
	}
	output, err := json.Marshal(response)
	if err == nil {
		w.Write(output)
	}
}

func main() {
	http.Handle("/securecookie-00/BenchmarkTest00403", &BenchmarkTest00403{})
	http.ListenAndServe(":8080", nil)
}
