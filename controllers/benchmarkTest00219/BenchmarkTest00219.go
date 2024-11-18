package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-sec-code/models"
	"log"
	"net/http"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", source)
	orm.Debug = true
}

type BenchmarkTest00219 struct{}

func (b *BenchmarkTest00219) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	db, err := sql.Open("mysql", source)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("select * from user where id=%s", id)
	user := models.User{}
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	output, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(output)
}

func main() {
	http.Handle("/pathtraver-00/BenchmarkTest00219", &BenchmarkTest00219{})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
