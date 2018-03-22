package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	testRouter := SetupRouter()

	req, err := http.NewRequest("GET", "/api/v1/users/1", nil)
	if err != nil {
		t.Errorf("Get hearteat failed with error %d.", err)
	}

	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)

	if resp.Code != 200 {
		t.Errorf("/api/v1/users failed with error code %d.", resp.Code)
	}
}

func TestGetUsers(t *testing.T) {
	gin.SetMode(gin.TestMode)
	testRouter := SetupRouter()

	req, err := http.NewRequest("GET", "/api/v1/users", nil)
	if err != nil {
		t.Errorf("Get hearteat failed with error %d.", err)
	}

	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)

	if resp.Code != 200 {
		t.Errorf("/api/v1/users failed with error code %d.", resp.Code)
	}
}

func TestPostUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	testRouter := SetupRouter()

	body := bytes.NewBuffer([]byte("{\"user_status\": \"83\", \"user_name\": \"100\"}"))

	req, err := http.NewRequest("POST", "/api/v1/users", body)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Errorf("Post hearteat failed with error %d.", err)
	}

	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)

	if resp.Code != 201 {
		t.Errorf("/api/v1/users failed with error code %d.", resp.Code)
	}
}

func TestPutUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	testRouter := SetupRouter()

	body := bytes.NewBuffer([]byte("{\"user_status\": \"83\", \"user_name\": \"100\"}"))

	req, err := http.NewRequest("PUT", "/api/v1/users/1", body)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Errorf("Put hearteat failed with error %d.", err)
	}

	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)

	if resp.Code != 200 {
		t.Errorf("/api/v1/users failed with error code %d.", resp.Code)
	}
}
