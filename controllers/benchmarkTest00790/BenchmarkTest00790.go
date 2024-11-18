package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00790 struct {
}

func (bt *BenchmarkTest00790) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest00790")

	encodedParam := url.QueryEscape(param)
	hashAndStore(encodedParam, w)
}

func hashAndStore(param string, w http.ResponseWriter) {
	// Assume hashing logic and storing in a file
	file, err := os.OpenFile("passwordFile.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		http.Error(w, "Could not open file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Simulating hash storage
	_, err = file.WriteString(fmt.Sprintf("hash_value=%s\n", param))
	if err != nil {
		http.Error(w, "Could not write to file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Sensitive value '%s' hashed and stored<br/>", param)
}

func main() {
	http.Handle("/hash-00/BenchmarkTest00790", &BenchmarkTest00790{})
	http.ListenAndServe(":8080", nil)
}
