package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00071 struct{}

func (b *BenchmarkTest00071) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	http.SetCookie(w, &http.Cookie{
		Name:   "BenchmarkTest00071",
		Value:  "someSecret",
		MaxAge: 60 * 3,
		Secure: true,
		Path:   r.RequestURI,
		Domain: r.URL.Hostname(),
	})

	http.ServeFile(w, r, "hash-00/BenchmarkTest00071.html")
}

func (b *BenchmarkTest00071) handlePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	theCookies := r.Cookies()
	param := "noCookieValueSupplied"
	for _, cookie := range theCookies {
		if cookie.Name == "BenchmarkTest00071" {
			param, _ = url.QueryUnescape(cookie.Value)
			break
		}
	}

	var bar string
	num := 196
	if (500/42)+num > 200 {
		bar = param
	} else {
		bar = "This should never happen"
	}

	hashFile, err := os.OpenFile("passwordFile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(w, "Could not open hash file", http.StatusInternalServerError)
		return
	}
	defer hashFile.Close()

	// Dummy hashing logic for demonstration
	hashValue := fmt.Sprintf("%x", bar)

	if _, err := hashFile.WriteString("hash_value=" + hashValue + "\n"); err != nil {
		http.Error(w, "Could not write to hash file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Sensitive value '%s' hashed and stored<br/>", bar)
	fmt.Fprintln(w, "Hash Test executed")
}

func main() {
	benchmark := &BenchmarkTest00071{}
	http.Handle("/", benchmark)
	http.HandleFunc("/hash-00/BenchmarkTest00071", benchmark.handlePost)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server failed:", err)
	}
}
