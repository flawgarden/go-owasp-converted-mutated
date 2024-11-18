package controllers

import (
	"fmt"
	"html"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest struct{}

func (b *BenchmarkTest) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")

	id := r.URL.Query().Get("BenchmarkTest01742")

	bar := b.doSomething(r, id)

	result, err := encrypt(bar)
	if err != nil {
		http.Error(w, "Error executing crypto", http.StatusInternalServerError)
		return
	}

	file, err := os.OpenFile("passwordFile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(w, "Error opening file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	if _, err := file.WriteString("secret_value=" + result + "\n"); err != nil {
		http.Error(w, "Error writing to file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Sensitive value: '%s' encrypted and stored<br/>", sanitizeInput(bar))
}

func (b *BenchmarkTest) doSomething(r *http.Request, param string) string {
	return sanitizeInput(param)
}

func sanitizeInput(input string) string {
	return html.EscapeString(input)
}

func encrypt(input string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(input), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
