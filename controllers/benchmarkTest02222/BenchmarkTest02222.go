package controllers

import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02222 struct{}

func (b *BenchmarkTest02222) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest02222")
	bar := doSomething(param)

	w.Header().Set("X-XSS-Protection", "0")
	_, err := fmt.Fprintf(w, bar)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func doSomething(param string) string {
	var bar string
	guess := "ABC"
	switchTarget := guess[1] // condition 'B', which is safe

	switch switchTarget {
	case 'A':
		bar = param
	case 'B':
		bar = "bob"
	case 'C', 'D':
		bar = param
	default:
		bar = "bob's your uncle"
	}

	return bar
}

func main() {
	http.Handle("/xss-04/BenchmarkTest02222", &BenchmarkTest02222{})
	fmt.Println("Server started at :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
