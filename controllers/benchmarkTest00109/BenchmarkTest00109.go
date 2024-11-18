package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"go-sec-code/models"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00109 struct{}

func (b *BenchmarkTest00109) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.SetCookie(w, &http.Cookie{
			Name:   "BenchmarkTest00109",
			Value:  "bar",
			Path:   r.RequestURI,
			MaxAge: 60 * 3,
			Secure: true,
		})
		http.ServeFile(w, r, "sqli-00/BenchmarkTest00109.html")
		return
	}

	if r.Method == http.MethodPost {
		cookies := r.Cookies()
		param := "noCookieValueSupplied"
		for _, cookie := range cookies {
			if cookie.Name == "BenchmarkTest00109" {
				param = cookie.Value
				break
			}
		}

		bar := ""
		if param != "" {
			decodedParam, err := url.QueryUnescape(param)
			if err == nil {
				bar = string(decodedParam)
			}
		}

		sqlStr := fmt.Sprintf("SELECT * FROM USERS WHERE USERNAME='foo' AND PASSWORD='%s'", bar)

		db, err := sql.Open("mysql", source)
		if err != nil {
			http.Error(w, "Error processing request.", http.StatusInternalServerError)
			return
		}
		defer db.Close()

		var user models.User
		err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
		if err != nil {
			http.Error(w, "Error processing request.", http.StatusInternalServerError)
			return
		}

		output, err := json.Marshal(user)
		if err == nil {
			w.Header().Set("Content-Type", "application/json")
			w.Write(output)
		}
	}
}
