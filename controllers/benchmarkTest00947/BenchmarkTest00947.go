package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00947 struct {
}

func (b *BenchmarkTest00947) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.SetCookie(w, &http.Cookie{
			Name:   "BenchmarkTest00947",
			Value:  "Ms+Bar",
			Path:   r.RequestURI,
			MaxAge: 60 * 3,
		})
		http.ServeFile(w, r, "ldapi-00/BenchmarkTest00947.html")
		return
	}

	if r.Method == http.MethodPost {
		var param = "noCookieValueSupplied"
		cookies := r.Cookies()
		for _, cookie := range cookies {
			if cookie.Name == "BenchmarkTest00947" {
				param = cookie.Value
				break
			}
		}

		bar := new(Test).doSomething(r, param)

		db, err := sql.Open("mysql", source)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer db.Close()

		filter := fmt.Sprintf("(&(objectclass=person))(|(uid=%s)(street={0}))", bar)
		filters := []string{"The streetz 4 Ms bar"}

		// LDAP search logic is omitted for brevity
		// Assume query execution and result processing here

		_, err = db.Exec(filter, filters)
		if err != nil {
			http.Error(w, "Query failed: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

type Test struct {
}

func (t *Test) doSomething(r *http.Request, param string) string {
	// Implement required processing
	return "processed_" + param
}

func main() {
	http.Handle("/ldapi-00/BenchmarkTest00947", &BenchmarkTest00947{})
	http.ListenAndServe(":8080", nil)
}
