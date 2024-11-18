package controllers

import (
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01534 struct{}

func (bt *BenchmarkTest01534) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.URL.Query().Get("BenchmarkTest01534")
	if param == "" {
		param = ""
	}

	secureRandomGenerator, err := os.Open("/dev/urandom")
	if err != nil {
		http.Error(w, "Error generating random bytes", http.StatusInternalServerError)
		return
	}
	defer secureRandomGenerator.Close()

	randomBytes := make([]byte, 40)
	if _, err := secureRandomGenerator.Read(randomBytes); err != nil {
		http.Error(w, "Error reading random bytes", http.StatusInternalServerError)
		return
	}

	rememberMeKey := fmt.Sprintf("%x", randomBytes)

	user := "SafeByron"
	testCaseNumber := "01534"
	user += testCaseNumber
	cookieName := "rememberMe" + testCaseNumber

	cookie, err := r.Cookie(cookieName)
	foundUser := false
	if err == nil {
		if cookie.Value == "" {
			foundUser = false
		} else {
			foundUser = true
		}
	}

	if foundUser {
		fmt.Fprintf(w, "Welcome back: %s<br/>", user)
	} else {
		http.SetCookie(w, &http.Cookie{
			Name:     cookieName,
			Value:    rememberMeKey,
			Secure:   true,
			HttpOnly: true,
			Path:     r.RequestURI,
		})
		fmt.Fprintf(w, "%s has been remembered with cookie: %s whose value is: %s<br/>", user, cookieName, rememberMeKey)
	}

	fmt.Fprint(w, "Randomness Test executed")
}

func (bt *BenchmarkTest01534) doSomething(param string) string {
	num := 106
	if (7*42)-num > 200 {
		return "This should never happen"
	}
	return param
}

func main() {
	http.Handle("/weakrand-03/BenchmarkTest01534", &BenchmarkTest01534{})
	http.ListenAndServe(":8080", nil)
}
