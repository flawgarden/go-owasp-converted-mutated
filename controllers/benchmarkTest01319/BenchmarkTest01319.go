package controllers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"go-sec-code/models"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01319 struct{}

func (b *BenchmarkTest01319) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		b.doPost(w, r)
		return
	}
	b.doPost(w, r)
}

func (b *BenchmarkTest01319) doPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest01319")
	bar := b.doSomething(param)

	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		http.Error(w, "Error generating key", http.StatusInternalServerError)
		return
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		http.Error(w, "Error creating cipher", http.StatusInternalServerError)
		return
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		http.Error(w, "Error creating GCM", http.StatusInternalServerError)
		return
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		http.Error(w, "Error generating nonce", http.StatusInternalServerError)
		return
	}

	ciphertext := gcm.Seal(nonce, nonce, []byte(bar), nil)

	encoded := base64.StdEncoding.EncodeToString(ciphertext)
	file, _ := json.Marshal(map[string]string{"secret_value": encoded})

	w.Header().Set("Content-Type", "application/json")
	w.Write(file)
}

func (b *BenchmarkTest01319) doSomething(param string) string {
	id, _ := strconv.Atoi(param)
	user := models.User{}
	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id=%d", id)
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		panic(err)
	}
	userData, _ := json.Marshal(user)
	return string(userData)
}
