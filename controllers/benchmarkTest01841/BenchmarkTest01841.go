package controllers

import (
	"net/http"
	"os"
)

type BenchmarkTest01841 struct{}

func (b *BenchmarkTest01841) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.SetCookie(w, &http.Cookie{
			Name:   "BenchmarkTest01841",
			Value:  "FileName",
			MaxAge: 60 * 3,
			Secure: true,
			Path:   r.URL.Path,
			Domain: r.URL.Hostname(),
		})
		http.ServeFile(w, r, "pathtraver-02/BenchmarkTest01841.html")
		return
	}

	if r.Method == http.MethodPost {
		cookie, err := r.Cookie("BenchmarkTest01841")
		param := "noCookieValueSupplied"
		if err == nil {
			param = cookie.Value
		}

		bar := doSomething(param)

		fileName := "path/to/testfiles/" + bar

		f, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0755)
		if err == nil {
			defer f.Close()
			_, _ = w.Write([]byte("Now ready to write to file: " + fileName))
		} else {
			http.Error(w, "Couldn't open FileOutputStream on file: '"+fileName+"'", http.StatusInternalServerError)
		}
	}
}

func doSomething(param string) string {
	a14546 := param
	b14546 := a14546 + " SafeStuff"
	b14546 = b14546[:len(b14546)-len("Chars")] + "Chars"
	map14546 := map[string]interface{}{"key14546": b14546}
	c14546 := map14546["key14546"].(string)
	d14546 := c14546[:len(c14546)-1]
	e14546 := string([]byte(d14546)) // B64 encoding omitted for brevity
	f14546 := e14546[:len(e14546)-len(" ")]
	bar := f14546 // Reflection omitted for brevity

	return bar
}
