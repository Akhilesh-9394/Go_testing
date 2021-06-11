package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestSignin(t *testing.T) {
	r := gin.New()
	r.POST("/signin", Signin)

	req, err := http.NewRequest("POST", "/signin", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestHome(t *testing.T) {
	r := gin.New()
	r.POST("/", Home)

	req, err := http.NewRequest("POST", "/signin", nil)
	cookie := http.Cookie{
		Name:    "token",
		Value:   tokenstring,
		Expires: expirationTime,
	}
	req.AddCookie(&cookie)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}
