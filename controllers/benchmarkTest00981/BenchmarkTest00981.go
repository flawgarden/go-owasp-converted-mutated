package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"go-sec-code/models"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00981Handler struct {
}

func (h *BenchmarkTest00981Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.SetCookie(w, &http.Cookie{
			Name:   "BenchmarkTest00981",
			Value:  ".",
			Path:   r.RequestURI,
			Secure: true,
			MaxAge: 60 * 3,
		})

		http.ServeFile(w, r, "cmdi-01/BenchmarkTest00981.html")
		return
	}

	if r.Method == http.MethodPost {
		cookie, err := r.Cookie("BenchmarkTest00981")
		param := "noCookieValueSupplied"
		if err == nil {
			param = cookie.Value
		}

		bar := doSomething(param)

		db, err := sql.Open("mysql", source)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer db.Close()

		sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id = %s", bar)
		user := models.User{}
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
		w.Header().Set("Content-Type", "application/json")
		w.Write(output)
		return
	}
}

func doSomething(param string) string {
	bar := param
	num := 106

	if (7*42)-num > 200 {
		bar = "This should never happen"
	}
	return bar
}
