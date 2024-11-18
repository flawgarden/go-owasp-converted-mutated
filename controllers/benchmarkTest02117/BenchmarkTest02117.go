package controllers

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"net/http"
	"strconv"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var users = map[int]User{
	1: {Id: 1, Username: "user1", Password: "pass1"},
	2: {Id: 2, Username: "user2", Password: "pass2"},
}

type BenchmarkTest02117 struct{}

func (b *BenchmarkTest02117) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	idStr := r.FormValue("BenchmarkTest02117")
	if idStr == "" {
		idStr = "0"
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	user, exists := users[id]
	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	output, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Error marshaling user", http.StatusInternalServerError)
		return
	}

	cookieName := "rememberMeBenchmarkTest02117"
	rememberMeKey := fmt.Sprintf("%f", rand.Float64())[2:]
	cookie := http.Cookie{
		Name:     cookieName,
		Value:    rememberMeKey,
		Secure:   true,
		HttpOnly: true,
		Path:     r.URL.Path,
	}
	http.SetCookie(w, &cookie)

	fmt.Fprintf(w, "User found: %s\n", string(output))
}

func main() {
	http.Handle("/weakrand-04/BenchmarkTest02117", &BenchmarkTest02117{})
	http.ListenAndServe(":8080", nil)
}
