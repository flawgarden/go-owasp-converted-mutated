package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-sec-code/models"
	"net/http"
	"strings"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01021 struct {
}

func (b *BenchmarkTest01021) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var param string
	if r.Header.Get("BenchmarkTest01021") != "" {
		param = r.Header.Get("BenchmarkTest01021")
	}
	param = strings.TrimSpace(param)

	bar := b.doSomething(param)

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id='%s'", bar)
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
}

func (b *BenchmarkTest01021) doSomething(param string) string {
	bar := "safe!"
	mapData := make(map[string]interface{})
	mapData["keyA-16616"] = "a_Value"
	mapData["keyB-16616"] = param
	bar = mapData["keyB-16616"].(string)
	bar = mapData["keyA-16616"].(string)
	return bar
}
