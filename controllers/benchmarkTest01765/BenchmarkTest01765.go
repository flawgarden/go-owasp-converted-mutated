package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01765 struct{}

func (b *BenchmarkTest01765) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		b.doPost(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (b *BenchmarkTest01765) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.URL.Query().Get("BenchmarkTest01765")
	bar := new(Test).doSomething(r, param)

	algorithm := "SHA512" // Could be loaded from properties file if needed.

	input := []byte(bar)
	hash := fmt.Sprintf("%x", CalculateHash(input, algorithm))

	file, err := os.OpenFile("passwordFile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(w, "Unable to open file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	if _, err := file.WriteString("hash_value=" + hash + "\n"); err != nil {
		http.Error(w, "Unable to write to file", http.StatusInternalServerError)
		return
	}

	_, _ = w.Write([]byte("Sensitive value '" + htmlEscape(string(input)) + "' hashed and stored<br/>"))
}

func CalculateHash(input []byte, algorithm string) []byte {
	// Implement hash calculation logic based on the algorithm
	return input // Placeholder
}

func htmlEscape(input string) string {
	return strings.ReplaceAll(strings.ReplaceAll(input, "&", "&amp;"), "<", "&lt;")
}

type Test struct{}

func (t *Test) doSomething(r *http.Request, param string) string {
	return param // Placeholder
}

func main() {
	http.Handle("/hash-02/BenchmarkTest01765", &BenchmarkTest01765{})
	http.ListenAndServe(":8080", nil)
}
