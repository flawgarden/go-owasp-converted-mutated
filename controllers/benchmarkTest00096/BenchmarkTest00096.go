package controllers

import (
	"math/rand"
	"net/http"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type BenchmarkTest00096 struct{}

func (b *BenchmarkTest00096) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	http.SetCookie(w, &http.Cookie{
		Name:   "BenchmarkTest00096",
		Value:  "whatever",
		Path:   r.URL.Path,
		Secure: true,
		MaxAge: 60 * 3,
	})

	http.ServeFile(w, r, "weakrand-00/BenchmarkTest00096.html")
}

func main() {
	http.Handle("/weakrand-00/BenchmarkTest00096", &BenchmarkTest00096{})
	http.ListenAndServe(":8080", nil)
}
