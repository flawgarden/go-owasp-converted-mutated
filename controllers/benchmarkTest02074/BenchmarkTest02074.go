package controllers

import (
	"fmt"
	"net/http"
	"os"
)

type BenchmarkTest02074 struct{}

func (b *BenchmarkTest02074) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var param string
	if headers, ok := r.Header["BenchmarkTest02074"]; ok && len(headers) > 0 {
		param = headers[0]
	}
	param = decode(param)

	rand, err := generateSecureRandom()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	rememberMeKey := fmt.Sprintf("%f", rand)[2:]
	user := "SafeDonna"
	testCaseNumber := "02074"
	user += testCaseNumber
	cookieName := "rememberMe" + testCaseNumber
	foundUser := false

	cookies := r.Cookies()
	for _, cookie := range cookies {
		if cookieName == cookie.Name && cookie.Value == r.Context().Value(cookieName).(string) {
			foundUser = true
			break
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
			Path:     r.URL.Path,
		})
		r.Context().Value(cookieName).(http.ResponseWriter).Write([]byte(rememberMeKey))
		fmt.Fprintf(w, "%s has been remembered with cookie: %s whose value is: %s<br/>", user, cookieName, rememberMeKey)
	}
	fmt.Fprintln(w, "Weak Randomness Test executed")
}

func decode(param string) string {
	return param // Decode logic implementation
}

func doSomething(param string) string {
	return param + "_SafeStuff"
}

func generateSecureRandom() (float64, error) {
	// Secure random generation implementation
	return 0.0, nil
}

func main() {
	http.Handle("/weakrand-04/BenchmarkTest02074", &BenchmarkTest02074{})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server Error:", err)
		os.Exit(1)
	}
}
