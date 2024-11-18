package controllers

import (
	"fmt"
	"net/http"
)

type BenchmarkTest00542 struct{}

func (b *BenchmarkTest00542) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		b.doPost(w, r)
		return
	}
	b.doPost(w, r)
}

func (b *BenchmarkTest00542) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	flag := true
	names := r.URL.Query()

	for name, values := range names {
		if flag {
			for _, value := range values {
				if value == "BenchmarkTest00542" {
					param = name
					flag = false
					break
				}
			}
		}
	}

	var bar string
	guess := "ABC"
	switchTarget := guess[2]

	switch switchTarget {
	case 'A':
		bar = param
	case 'B':
		bar = "bobs_your_uncle"
	case 'C', 'D':
		bar = param
	default:
		bar = "bobs_your_uncle"
	}

	w.Header().Set("X-XSS-Protection", "0")
	if _, err := fmt.Fprintf(w, bar); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.Handle("/xss-00/BenchmarkTest00542", &BenchmarkTest00542{})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
