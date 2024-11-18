package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02211Controller struct {
	http.ResponseWriter
	*http.Request
}

func (c *BenchmarkTest02211Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.ResponseWriter = w
	c.Request = r
	if r.Method == http.MethodGet {
		c.doPost()
	} else {
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
	}
}

func (c *BenchmarkTest02211Controller) doPost() {
	c.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.Request.URL.Query().Get("BenchmarkTest02211")
	bar := doSomething(param)

	hash, err := bcrypt.GenerateFromPassword([]byte(bar), bcrypt.DefaultCost)
	if err != nil {
		http.Error(c.ResponseWriter, "Hashing error", http.StatusInternalServerError)
		return
	}

	file, err := os.OpenFile("passwordFile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(c.ResponseWriter, "File error", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	if _, err := file.WriteString("hash_value=" + string(hash) + "\n"); err != nil {
		http.Error(c.ResponseWriter, "File write error", http.StatusInternalServerError)
		return
	}
	_, _ = fmt.Fprintf(c.ResponseWriter, "Sensitive value '%s' hashed and stored<br/>", htmlEscape(bar))
}

func doSomething(param string) string {
	return htmlEscape(param)
}

func htmlEscape(input string) string {
	return strconv.QuoteToASCII(input)
}
