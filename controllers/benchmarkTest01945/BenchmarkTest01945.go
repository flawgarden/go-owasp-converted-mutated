package controllers

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
)

type BenchmarkTest01945 struct{}

func (b *BenchmarkTest01945) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.Header.Add("Content-Type", "text/html;charset=UTF-8")

	param := r.Header.Get("BenchmarkTest01945")
	param, _ = url.QueryUnescape(param)

	secureRandom := make([]byte, 40)
	if _, err := rand.Read(secureRandom); err != nil {
		http.Error(w, "Randomness error", http.StatusInternalServerError)
		return
	}

	rememberMeKey := base64.StdEncoding.EncodeToString(secureRandom)
	user := "SafeByron"
	fullClassName := fmt.Sprintf("%T", b)
	testCaseNumber := fullClassName[len(fullClassName)-len("BenchmarkTest01945"):]

	user += testCaseNumber
	cookieName := "rememberMe" + testCaseNumber

	cookies := r.Cookies()
	foundUser := false
	for _, cookie := range cookies {
		if cookieName == cookie.Name && cookie.Value == r.Context().Value(cookieName) {
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
			Path:     r.RequestURI,
		})
		r.Context().Value(cookieName) // store in context (for demonstration)
		fmt.Fprintf(w, "%s has been remembered with cookie: %s whose value is: %s<br/>", user, cookieName, rememberMeKey)
	}

	fmt.Fprintln(w, "Randomness Test executed")
}

func doSomething(r *http.Request, param string) string {
	bar := ""
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the param value
	}
	return bar
}

func main() {
	http.Handle("/weakrand-04/BenchmarkTest01945", &BenchmarkTest01945{})
	http.ListenAndServe(":8080", nil)
}
