package controllers

import (
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"database/sql"
	"encoding/json"
	"fmt"
	"go-sec-code/models"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type SqlInjectionVuln1Controller struct{}

func (c *SqlInjectionVuln1Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		c.doGet(w, r)
	} else if r.Method == http.MethodPost {
		c.doPost(w, r)
	}
}

func (c *SqlInjectionVuln1Controller) doGet(w http.ResponseWriter, r *http.Request) {
	c.doPost(w, r)
}

func (c *SqlInjectionVuln1Controller) doPost(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("select * from user where id=%s", id)
	user := models.User{}
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		panic(err)
	}

	output, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	// DES encryption
	block, err := des.NewCipher([]byte("12345678"))
	if err != nil {
		log.Fatal(err)
	}

	ciphertext := make([]byte, des.BlockSize)
	iv := make([]byte, des.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		log.Fatal(err)
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, output)

	err = ioutil.WriteFile("passwordFile.txt", ciphertext, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte("Sensitive value encrypted and stored"))
}
