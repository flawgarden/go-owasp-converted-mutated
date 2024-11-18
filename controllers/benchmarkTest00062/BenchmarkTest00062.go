package controllers

import (
	"net/http"
	"os"
	"path/filepath"
)

type BenchmarkTest00062 struct{}

func (b *BenchmarkTest00062) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.SetCookie(w, &http.Cookie{
			Name:   "BenchmarkTest00062",
			Value:  "FileName",
			Path:   r.URL.Path,
			MaxAge: 180,
			Secure: true,
		})
		http.ServeFile(w, r, "pathtraver-00/BenchmarkTest00062.html")
	} else if r.Method == http.MethodPost {
		var param string = "noCookieValueSupplied"
		cookies := r.Cookies()
		for _, cookie := range cookies {
			if cookie.Name == "BenchmarkTest00062" {
				param = cookie.Value
				break
			}
		}

		bar := "safe!"
		map77232 := map[string]string{
			"keyA-77232": "a-Value",
			"keyB-77232": param,
			"keyC":       "another-Value",
		}
		bar = map77232["keyB-77232"]

		fileName := filepath.Join("testfiles", bar)
		fis, err := os.Open(fileName)
		if err != nil {
			http.Error(w, "Problem getting FileInputStream: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer fis.Close()

		b := make([]byte, 1000)
		size, _ := fis.Read(b)
		w.Write([]byte("The beginning of file: '" + fileName + "' is:\n\n" + string(b[:size])))
	}
}
