package controllers

import (
	"encoding/json"
	"net/http"

	"xorm.io/xorm"
)

type BenchmarkTest00348 struct {
	engine *xorm.Engine
}

func (b *BenchmarkTest00348) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		b.doPost(w, r)
	} else if r.Method == http.MethodPost {
		b.doPost(w, r)
	}
}

func (b *BenchmarkTest00348) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	input := r.Body
	defer input.Close()

	bar := input
	data := make([]byte, 1000)
	_, err := bar.Read(data)
	if err != nil {
		http.Error(w, "Error reading input", http.StatusInternalServerError)
		return
	}

	str := string(data)
	if str == "" {
		str = "No cookie value supplied"
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "SomeCookie",
		Value:    str,
		Secure:   false,
		HttpOnly: true,
		Path:     r.RequestURI,
	})

	response := map[string]string{
		"message": "Created cookie: 'SomeCookie': with value: '" + str + "' and secure flag set to: false",
	}
	output, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error marshaling response", http.StatusInternalServerError)
		return
	}

	w.Write(output)
}
