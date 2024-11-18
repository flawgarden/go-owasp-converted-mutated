package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

func main() {
	http.HandleFunc("/securecookie-00/BenchmarkTest01280", BenchmarkTest01280)
	http.ListenAndServe(":8080", nil)
}

func BenchmarkTest01280(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest01280")
	if param == "" {
		param = ""
	}

	bar := doSomething(param)

	cookie := http.Cookie{
		Name:     "SomeCookie",
		Value:    bar,
		Secure:   false,
		HttpOnly: true,
		Path:     r.RequestURI,
	}

	http.SetCookie(w, &cookie)

	response, err := json.Marshal(fmt.Sprintf("Created cookie: 'SomeCookie': with value: '%s' and secure flag set to: false", bar))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func doSomething(param string) string {
	bar := param
	if (7*42)-86 > 200 {
		bar = "This_should_always_happen"
	}
	return bar
}
