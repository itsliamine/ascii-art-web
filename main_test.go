package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestHomeHandler verifies that the home handler returns a status code of 200 (OK).
func TestHomeHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(homeHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

// TestNotFoundHandler verifies that a non-existent route returns a status code of 404 (Not Found).
func TestNotFoundHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/nonexistent", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(homeHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}
}

// TestPostHandler verifies that a valid POST request to the /ascii endpoint returns a status code of 302 (Found).
func TestPostHandler(t *testing.T) {
	form := strings.NewReader("input=hello&banner=standard")
	req, err := http.NewRequest("POST", "/ascii", form)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(postHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusFound)
	}
}

// TestBadRequestHandler verifies that a GET request to the /ascii endpoint returns a status code of 400 (Bad Request).
func TestBadRequestHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/ascii", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(postHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

// TestInternalServerErrorHandler verifies that a malformed POST request to the /ascii endpoint returns a status code of 500 (Internal Server Error).
func TestInternalServerErrorHandler(t *testing.T) {
	req, err := http.NewRequest("POST", "/ascii", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(postHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusInternalServerError)
	}
}
