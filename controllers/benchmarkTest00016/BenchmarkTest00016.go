package controllers

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

type BenchmarkTest00016 struct{}

func (b *BenchmarkTest00016) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		b.doPost(w, r)
	} else if r.Method == http.MethodPost {
		b.doPost(w, r)
	}
}

func (b *BenchmarkTest00016) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	headers := r.Header["BenchmarkTest00016"]

	if len(headers) > 0 {
		param = headers[0]
	}

	param, _ = url.QueryUnescape(param)

	var input string
	if strings.TrimSpace(param) == "" {
		input = "No cookie value supplied"
	} else {
		input = param
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "SomeCookie",
		Value:    input,
		Secure:   true,
		HttpOnly: true,
		Path:     r.RequestURI,
	})

	response := map[string]string{"message": "Created cookie: 'SomeCookie' with value: '" + input + "' and secure flag set to: true"}
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.Handle("/securecookie-00/BenchmarkTest00016", &BenchmarkTest00016{})
	http.ListenAndServe(":8080", nil)
}
