package controllers

import (
	"net/http"
	"net/url"
	"os"
)

type BenchmarkTest00954 struct{}

func (b *BenchmarkTest00954) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	userCookie := http.Cookie{
		Name:   "BenchmarkTest00954",
		Value:  "FileName",
		Path:   r.RequestURI,
		MaxAge: 180,
		Secure: true,
	}
	http.SetCookie(w, &userCookie)

	http.ServeFile(w, r, "pathtraver-01/BenchmarkTest00954.html")
}

func (b *BenchmarkTest00954) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	cookies := r.Cookies()

	param := "noCookieValueSupplied"
	for _, cookie := range cookies {
		if cookie.Name == "BenchmarkTest00954" {
			decodedValue, _ := url.QueryUnescape(cookie.Value)
			param = decodedValue
			break
		}
	}

	bar := b.doSomething(r, param)
	fileName := "TESTFILES_DIR/" + bar

	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		http.Error(w, "Couldn't open FileOutputStream on file: '"+fileName+"'", http.StatusInternalServerError)
		return
	}
	defer f.Close()

	w.Write([]byte("Now ready to write to file: " + fileName))
}

func (b *BenchmarkTest00954) doSomething(r *http.Request, param string) string {
	bar := "safe!"
	collection := make(map[string]interface{})
	collection["keyA-9749"] = "a_Value"
	collection["keyB-9749"] = param
	collection["keyC"] = "another_Value"
	bar = collection["keyB-9749"].(string)
	bar = collection["keyA-9749"].(string)

	return bar
}

func main() {
	test := &BenchmarkTest00954{}
	http.Handle("/pathtraver-01/BenchmarkTest00954", test)
	http.ListenAndServe(":8080", nil)
}
