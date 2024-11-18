package controllers

import (
	"encoding/json"
	"net/http"
)

type BenchmarkTest02247 struct{}

func (b *BenchmarkTest02247) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := r.URL.Query().Get("BenchmarkTest02247")
	bar := doSomething(param)

	cookie := http.Cookie{
		Name:     "SomeCookie",
		Value:    bar,
		Secure:   true,
		HttpOnly: true,
		Path:     r.RequestURI,
	}
	http.SetCookie(w, &cookie)

	output := struct {
		Message string `json:"message"`
		Cookie  string `json:"cookie"`
	}{
		Message: "Created cookie: 'SomeCookie': with value: '" + bar + "' and secure flag set to: true",
		Cookie:  bar,
	}

	json.NewEncoder(w).Encode(output)
}

func doSomething(param string) string {
	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = append(valuesList[:0], valuesList[1:]...) // remove the 1st safe value
		bar = valuesList[1]                                    // get the last 'safe' value
	}
	return bar
}
