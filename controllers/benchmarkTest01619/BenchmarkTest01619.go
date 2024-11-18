package controllers

import (
	"fmt"
	"html"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01619 struct{}

func (b *BenchmarkTest01619) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	values := r.Form["BenchmarkTest01619"]
	var param string
	if len(values) > 0 {
		param = values[0]
	} else {
		param = ""
	}

	bar := escapeHTML(param)

	r.Header.Set("userid", bar)
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	fmt.Fprintf(w, "Item: 'userid' with value: '%s' saved in session.", bar)
}

func escapeHTML(input string) string {
	return html.EscapeString(input)
}

func main() {
	http.Handle("/trustbound-01/BenchmarkTest01619", &BenchmarkTest01619{})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
