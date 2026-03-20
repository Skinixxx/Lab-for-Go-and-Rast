package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHTTPHandler_Success(t *testing.T) {
	req, _ := http.NewRequest("POST", "/integrate?func=sin&a=0&b=3.14159&n=10000", nil)
	w := httptest.NewRecorder()

	httpHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	if response["result"] == nil {
		t.Error("Expected result in response")
	}
}

func TestHTTPHandler_WrongMethod(t *testing.T) {
	req, _ := http.NewRequest("GET", "/integrate?func=sin&a=0&b=3.14159&n=10000", nil)
	w := httptest.NewRecorder()

	httpHandler(w, req)

	if w.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected status 405, got %d", w.Code)
	}
}

func TestHTTPHandler_Sin(t *testing.T) {
	req, _ := http.NewRequest("POST", "/integrate?func=sin&a=0&b=3.14159&n=100000", nil)
	w := httptest.NewRecorder()

	httpHandler(w, req)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	result := response["result"].(float64)
	if result < 1.99 || result > 2.01 {
		t.Errorf("Expected ~2.0, got %f", result)
	}
}

func TestHTTPHandler_Cos(t *testing.T) {
	req, _ := http.NewRequest("POST", "/integrate?func=cos&a=0&b=1.5708&n=50000", nil)
	w := httptest.NewRecorder()

	httpHandler(w, req)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	result := response["result"].(float64)
	if result < 0.99 || result > 1.01 {
		t.Errorf("Expected ~1.0, got %f", result)
	}
}

func TestHTTPHandler_XSquared(t *testing.T) {
	req, _ := http.NewRequest("POST", "/integrate?func=x^2&a=0&b=1&n=100000", nil)
	w := httptest.NewRecorder()

	httpHandler(w, req)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	result := response["result"].(float64)
	expected := 1.0 / 3.0
	if result < expected-0.001 || result > expected+0.001 {
		t.Errorf("Expected ~%f, got %f", expected, result)
	}
}

func TestHTTPHandler_Exp(t *testing.T) {
	req, _ := http.NewRequest("POST", "/integrate?func=exp&a=0&b=1&n=100000", nil)
	w := httptest.NewRecorder()

	httpHandler(w, req)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	result := response["result"].(float64)
	expected := 2.718281828459045 - 1
	if result < expected-0.001 || result > expected+0.001 {
		t.Errorf("Expected ~%f, got %f", expected, result)
	}
}

func TestHTTPHandler_X(t *testing.T) {
	req, _ := http.NewRequest("POST", "/integrate?func=x&a=0&b=1&n=100000", nil)
	w := httptest.NewRecorder()

	httpHandler(w, req)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	result := response["result"].(float64)
	if result < 0.49 || result > 0.51 {
		t.Errorf("Expected ~0.5, got %f", result)
	}
}

func TestHTTPHandler_MissingParams(t *testing.T) {
	req, _ := http.NewRequest("POST", "/integrate?func=sin", nil)
	w := httptest.NewRecorder()

	httpHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200 even with missing params, got %d", w.Code)
	}
}

func TestHTTPHandler_ContentType(t *testing.T) {
	req, _ := http.NewRequest("POST", "/integrate?func=sin&a=0&b=1&n=1000", nil)
	w := httptest.NewRecorder()

	httpHandler(w, req)

	contentType := w.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Expected Content-Type application/json, got %s", contentType)
	}
}

func BenchmarkHTTPHandler_Sin(b *testing.B) {
	reqBody := "/integrate?func=sin&a=0&b=3.14159&n=10000"
	req, _ := http.NewRequest("POST", reqBody, &bytes.Buffer{})
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		httpHandler(w, req)
	}
}
