package controllers

import (
	"fmt"
	"net/http"

	"xorm.io/xorm"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00058 struct {
	Db *xorm.Engine
}

func (b *BenchmarkTest00058) Get(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "BenchmarkTest00058",
		Value:  "someSecret",
		MaxAge: 180,
		Secure: true,
		Path:   r.URL.Path,
		Domain: r.URL.Host,
	})
	http.StripPrefix("/crypto-00/", http.FileServer(http.Dir("path/to/html"))).ServeHTTP(w, r)
}

func (b *BenchmarkTest00058) Post(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("BenchmarkTest00058")
	param := "noCookieValueSupplied"
	if err == nil {
		param = cookie.Value
	}

	bar := param
	result, err := encryptInput(bar)
	if err != nil {
		http.Error(w, "Problem executing crypto", http.StatusInternalServerError)
		return
	}

	err = storeEncryptedValue(result, bar)
	if err != nil {
		http.Error(w, "Failed to store value", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Sensitive value: '%s' encrypted and stored<br/>", bar)
}

func encryptInput(input string) ([]byte, error) {
	// TODO: Implement encryption logic here
	return nil, nil
}

func storeEncryptedValue(value []byte, original string) error {
	// TODO: Implement storage logic here
	return nil
}
