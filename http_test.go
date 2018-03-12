package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := ginEngine()
	req, _ := http.NewRequest("GET", "/", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, resp.Code, 200)
}

func TestURLNotFound(t *testing.T) {
	setup()
	gin.SetMode(gin.TestMode)
	router := ginEngine()
	req, _ := http.NewRequest("GET", "/foobar", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, resp.Code, 404)
}

func TestEmptyForm(t *testing.T) {
	setup()
	gin.SetMode(gin.TestMode)
	router := ginEngine()
	req, _ := http.NewRequest("POST", "/", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, resp.Code, 200)
	assert.Contains(t, resp.Body.String(), "invalid-feedback")
}

func TestBadFormData(t *testing.T) {
	setup()
	gin.SetMode(gin.TestMode)
	router := ginEngine()
	form := url.Values{}
	form.Add("url", "---")
	req, _ := http.NewRequest("POST", "/", nil)
	req.PostForm = form
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, resp.Code, 200)
	assert.Contains(t, resp.Body.String(), "invalid-feedback")
}

func TestFormCreateURL(t *testing.T) {
	setup()
	gin.SetMode(gin.TestMode)
	router := ginEngine()
	form := url.Values{}
	form.Add("url", "foo.com/?bar=2#goto")
	req, _ := http.NewRequest("POST", "/", nil)
	req.PostForm = form
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, resp.Code, 201)
	assert.NotContains(t, resp.Body.String(), "invalid-feedback")
	assert.Contains(t, resp.Body.String(), "Copy to clipboard")
}
