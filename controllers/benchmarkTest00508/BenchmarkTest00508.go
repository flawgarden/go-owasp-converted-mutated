package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00508 struct{}

func (b *BenchmarkTest00508) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("BenchmarkTest00508")
	bar := fmt.Sprintf("%s_SafeStuff", param)

	http.SetCookie(w, &http.Cookie{Name: "userid", Value: bar})

	output, err := json.Marshal(map[string]string{"userid": bar})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func main() {
	http.Handle("/trustbound-00/BenchmarkTest00508", &BenchmarkTest00508{})
	http.ListenAndServe(":8080", nil)
}
