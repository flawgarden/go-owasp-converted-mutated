package controllers

import (
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

type BenchmarkTest01836 struct {
	http.Handler
}

func (b *BenchmarkTest01836) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.SetCookie(w, &http.Cookie{
			Name:   "BenchmarkTest01836",
			Value:  "FileName",
			MaxAge: 60 * 3,
			Secure: true,
			Path:   r.URL.Path,
			Domain: r.Host,
		})
		http.ServeFile(w, r, "pathtraver-02/BenchmarkTest01836.html")
		return
	}

	if r.Method == http.MethodPost {
		cookies := r.Cookies()
		param := "noCookieValueSupplied"
		for _, cookie := range cookies {
			if cookie.Name == "BenchmarkTest01836" {
				decodedValue, _ := url.QueryUnescape(cookie.Value)
				param = decodedValue
				break
			}
		}

		bar := doSomething(r, param)
		startURIslashes := ""
		if os := os.Getenv("OS"); os != "" && os == "Windows_NT" {
			startURIslashes = "/"
		} else {
			startURIslashes = "//"
		}

		fileURI := "file:" + startURIslashes + filepath.Join("testfiles", bar)
		fileTarget := filepath.Clean(fileURI)
		w.Write([]byte("Access to file: '" + fileTarget + "' created.\n"))
		if _, err := os.Stat(fileTarget); err == nil {
			w.Write([]byte(" And file already exists.\n"))
		} else {
			w.Write([]byte(" But file doesn't exist yet.\n"))
		}
	}
}

func doSomething(r *http.Request, param string) string {
	bar := "safe!"
	map78713 := make(map[string]interface{})
	map78713["keyA-78713"] = "a-Value"
	map78713["keyB-78713"] = param
	map78713["keyC"] = "another-Value"
	bar = map78713["keyB-78713"].(string)
	return bar
}
