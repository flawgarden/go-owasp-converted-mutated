package controllers

import (
	"fmt"
	"net/http"
)

type BenchmarkTest00469 struct{}

func (bt *BenchmarkTest00469) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.Redirect(w, r, r.URL.Path, http.StatusSeeOther)
		return
	}
	r.ParseForm()

	var param string
	if values, ok := r.Form["BenchmarkTest00469"]; ok && len(values) > 0 {
		param = values[0]
	}

	var bar string
	num := 106
	if (7*18)+num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}

	w.Header().Set("X-XSS-Protection", "0")
	_, err := fmt.Fprintf(w, bar)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.Handle("/xss-00/BenchmarkTest00469", &BenchmarkTest00469{})
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
