package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00086 struct{}

func (b *BenchmarkTest00086) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	http.SetCookie(w, &http.Cookie{
		Name:    "BenchmarkTest00086",
		Value:   "whatever",
		Expires: time.Now().Add(3 * time.Minute),
		Secure:  true,
		Path:    r.URL.Path,
		Domain:  r.Host,
	})

	if r.Method == http.MethodPost {
		cookies := r.Cookies()
		param := "noCookieValueSupplied"
		for _, cookie := range cookies {
			if cookie.Name == "BenchmarkTest00086" {
				param, _ = url.QueryUnescape(cookie.Value)
				break
			}
		}

		num := 86
		bar := param
		if (7*42)-num > 200 {
			bar = "This_should_always_happen"
		}

		id := bar
		db, err := sql.Open("mysql", source)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer db.Close()

		sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id=%s", id)
		user := struct {
			Id       int
			Username string
			Password string
		}{}
		err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		output, err := json.Marshal(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(output)
		return
	}

	http.Redirect(w, r, "/weakrand-00/BenchmarkTest00086.html", http.StatusSeeOther)
}

func main() {
	http.Handle("/weakrand-00/BenchmarkTest00086", &BenchmarkTest00086{})
	http.ListenAndServe(":8080", nil)
}
