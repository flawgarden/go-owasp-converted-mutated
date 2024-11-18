package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-sec-code/models"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01501 struct{}

func (b *BenchmarkTest01501) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest01501")
	if param == "" {
		param = ""
	}

	bar := b.doSomething(param)

	output, err := b.queryLDAP(bar)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	w.Write([]byte(output))
}

func (b *BenchmarkTest01501) doSomething(param string) string {
	valuesList := []string{"safe", param, "moresafe"}
	valuesList = append(valuesList[:0], valuesList[1:]...) // remove the 1st safe value
	return valuesList[0]                                   // get the param value
}

func (b *BenchmarkTest01501) queryLDAP(bar string) (string, error) {
	db, err := sql.Open("mysql", source)
	if err != nil {
		return "", err
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE username='%s'", bar)
	user := models.User{}
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		return "", err
	}

	output, err := json.Marshal(user)
	if err != nil {
		return "", err
	}

	return string(output), nil
}
