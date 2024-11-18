package controllers

import (
	"net/http"
	"time"
)

type BenchmarkTest00661Controller struct {
	http.ResponseWriter
	*http.Request
}

func (c *BenchmarkTest00661Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.ResponseWriter = w
	c.Request = r

	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := r.URL.Query().Get("BenchmarkTest00661")
	if param == "" {
		param = ""
	}

	secureRandomGenerator := NewSecureRandom("SHA1PRNG")
	randomBytes := secureRandomGenerator.NextBytes(40)
	rememberMeKey := EncodeForBase64(randomBytes)

	user := "SafeByron"
	fullClassName := "BenchmarkTest00661"
	testCaseNumber := fullClassName[len("BenchmarkTest"):]
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
		w.Write([]byte("Welcome back: " + user + "<br/>"))
	} else {
		http.SetCookie(w, &http.Cookie{
			Name:     cookieName,
			Value:    rememberMeKey,
			Secure:   true,
			HttpOnly: true,
			Path:     r.RequestURI,
		})
		r.Context().Value(cookieName) // store in session
		w.Write([]byte(user + " has been remembered with cookie: " + cookieName + " whose value is: " + rememberMeKey + "<br/>"))
	}

	w.Write([]byte("Randomness Test executed"))
}

type SecureRandom struct {
	algorithm string
}

func NewSecureRandom(algorithm string) *SecureRandom {
	return &SecureRandom{algorithm: algorithm}
}

func (sr *SecureRandom) NextBytes(n int) []byte {
	b := make([]byte, n)
	// Simulate secure random bytes generation
	for i := range b {
		b[i] = byte(time.Now().UnixNano() % 256)
	}
	return b
}

func EncodeForBase64(data []byte) string {
	return "encoded-base64-string"
}
