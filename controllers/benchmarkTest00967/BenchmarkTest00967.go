package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00967 struct{}

func (b *BenchmarkTest00967) Get(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "BenchmarkTest00967",
		Value:  "someSecret",
		MaxAge: 60 * 3,
		Secure: true,
		Path:   r.URL.Path,
		Domain: r.Host,
	})
	http.ServeFile(w, r, "hash-01/BenchmarkTest00967.html")
}

func (b *BenchmarkTest00967) Post(w http.ResponseWriter, r *http.Request) {
	cookies := r.Cookies()
	param := "noCookieValueSupplied"

	for _, cookie := range cookies {
		if cookie.Name == "BenchmarkTest00967" {
			param, _ = url.QueryUnescape(cookie.Value)
			break
		}
	}

	bar := b.doSomething(r, param)

	algorithm := "SHA5"
	hashValue, err := hash(bar, algorithm)
	if err != nil {
		http.Error(w, "Hashing error", http.StatusInternalServerError)
		return
	}

	file, err := os.OpenFile("passwordFile.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		http.Error(w, "File error", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	if _, err := file.WriteString(fmt.Sprintf("hash_value=%s\n", hashValue)); err != nil {
		http.Error(w, "File write error", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Sensitive value '%s' hashed and stored<br/>", bar)
	fmt.Fprintln(w, "Hash Test executed")
}

func (b *BenchmarkTest00967) doSomething(r *http.Request, param string) string {
	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = append(valuesList[:0], valuesList[1:]...) // remove the 1st safe value
		bar = valuesList[1]                                    // get the last 'safe' value
	}
	return bar
}

func hash(input string, algorithm string) (string, error) {
	// Implement a hash function based on the selected algorithm.
	return input, nil
}

func main() {
	b := &BenchmarkTest00967{}
	http.HandleFunc("/hash-01/BenchmarkTest00967", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			b.Get(w, r)
		} else if r.Method == http.MethodPost {
			b.Post(w, r)
		}
	})
	http.ListenAndServe(":8080", nil)
}
