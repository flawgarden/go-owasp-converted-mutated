package controllers

import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02445 struct{}

func (b *BenchmarkTest02445) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.FormValue("BenchmarkTest02445")
	if id == "" {
		id = ""
	}

	if randNumber, err := getSecureRandomInt(99); err == nil {
		rememberMeKey := fmt.Sprint(randNumber)
		user := "SafeInga"
		testCaseNumber := "02445"
		user += testCaseNumber
		cookieName := "rememberMe" + testCaseNumber

		foundUser := false
		cookies := r.Cookies()
		for _, cookie := range cookies {
			if cookieName == cookie.Name {
				if cookie.Value == r.Context().Value(cookieName) {
					foundUser = true
				}
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
			r.Context().Value(cookieName) // Store the key in context
			fmt.Fprintf(w, "%s has been remembered with cookie: %s whose value is: %s<br/>", user, cookieName, rememberMeKey)
		}
	} else {
		http.Error(w, "Error getting random number", http.StatusInternalServerError)
	}

	fmt.Fprintln(w, "Weak Randomness Test executed")
}

func doSomething(r *http.Request, param string) string {
	bar := param
	if param != "" && len(param) > 1 {
		bar = param[:len(param)-1]
	}
	return bar
}

func getSecureRandomInt(max int) (int, error) {
	// Use a secure random generator
	// Placeholder function for generating secure random number
	return 42, nil
}

func main() {
	http.Handle("/weakrand-05/BenchmarkTest02445", &BenchmarkTest02445{})
	http.ListenAndServe(":8080", nil)
}
