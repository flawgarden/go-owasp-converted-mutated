package controllers

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func BenchmarkTest01838(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	userCookie := http.Cookie{
		Name:   "BenchmarkTest01838",
		Value:  "FileName",
		MaxAge: 180,
		Secure: true,
		Path:   r.RequestURI,
		Domain: r.Host,
	}
	http.SetCookie(w, &userCookie)
	http.ServeFile(w, r, "pathtraver-02/BenchmarkTest01838.html")
}

func TestPostRequest(t *testing.T) {
	req, err := http.NewRequest("POST", "/pathtraver-02/BenchmarkTest01838", nil)
	if err != nil {
		t.Fatal(err)
	}
	cookie := &http.Cookie{Name: "BenchmarkTest01838", Value: "testFile.txt"}
	req.AddCookie(cookie)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(BenchmarkTest01838)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "The beginning of file: '"
	fileName := "pathtraver-02/testFile.txt"
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		t.Errorf("File does not exist: %v", fileName)
		return
	}

	fis, err := os.Open(fileName)
	if err == nil {
		defer fis.Close()
		b := make([]byte, 1000)
		size, _ := fis.Read(b)
		if rr.Body.String() != expected+fileName+"' is:\n\n"+string(b[:size]) {
			t.Errorf("Unexpected body: got %v want %v", rr.Body.String(), expected+fileName+"' is:\n\n"+string(b[:size]))
		}
	} else {
		t.Errorf("Couldn't open FileInputStream on file: '%s'", fileName)
	}
}
