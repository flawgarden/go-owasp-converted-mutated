package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00056 struct {
}

func (bt *BenchmarkTest00056) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	cookie := &http.Cookie{
		Name:   "BenchmarkTest00056",
		Value:  "someSecret",
		MaxAge: 180, // Store cookie for 3 minutes
		Secure: true,
		Path:   r.URL.Path,
		Domain: r.URL.Hostname(),
	}
	http.SetCookie(w, cookie)

	if r.Method == http.MethodGet {
		http.ServeFile(w, r, "crypto-00/BenchmarkTest00056.html")
		return
	}

	param := "noCookieValueSupplied"
	cookies := r.Cookies()
	for _, c := range cookies {
		if c.Name == "BenchmarkTest00056" {
			param, _ = url.QueryUnescape(c.Value)
			break
		}
	}

	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[1]         // get the last 'safe' value
	}

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id=%s", bar)
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
}

func main() {
	http.Handle("/crypto-00/BenchmarkTest00056", &BenchmarkTest00056{})
	http.ListenAndServe(":8080", nil)
}
