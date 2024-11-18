package controllers

import (
	"fmt"
	"html"
	"net/http"
	"time"
)

type BenchmarkTest00718 struct{}

func (b *BenchmarkTest00718) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		b.doPost(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (b *BenchmarkTest00718) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	values := r.URL.Query()["BenchmarkTest00718"]
	var param string
	if len(values) > 0 {
		param = values[0]
	} else {
		param = ""
	}

	bar := html.EscapeString(param)

	w.Header().Set("X-XSS-Protection", "0")
	_, _ = fmt.Fprintf(w, "Formatted like: %s and %s.", "a", bar)
}

func main() {
	http.Handle("/xss-01/BenchmarkTest00718", &BenchmarkTest00718{})
	srv := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	srv.ListenAndServe()
}
