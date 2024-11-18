package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00986 struct{}

func (b *BenchmarkTest00986) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.SetCookie(w, &http.Cookie{
			Name:   "BenchmarkTest00986",
			Value:  "whatever",
			MaxAge: 60 * 3,
			Secure: true,
			Path:   r.RequestURI,
			Domain: r.Host,
		})
		http.ServeFile(w, r, "weakrand-02/BenchmarkTest00986.html")
	} else if r.Method == http.MethodPost {
		cookies := r.Cookies()
		param := "noCookieValueSupplied"
		for _, cookie := range cookies {
			if cookie.Name == "BenchmarkTest00986" {
				param, _ = url.QueryUnescape(cookie.Value)
				break
			}
		}

		bar := param
		if len(param) > 1 {
			bar = param[:len(param)-1]
		}

		db, err := sql.Open("mysql", source)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer db.Close()

		sqlStr := fmt.Sprintf("SELECT * FROM user WHERE username='%s'", bar)
		var user struct {
			Id       int
			Username string
			Password string
		}
		err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_, err = w.Write([]byte(fmt.Sprintf("User: %s", user.Username)))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func main() {
	http.Handle("/weakrand-02/BenchmarkTest00986", &BenchmarkTest00986{})
	http.ListenAndServe(":8080", nil)
}
