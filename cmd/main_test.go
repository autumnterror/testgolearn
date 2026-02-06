package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestStatusHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/status", nil)
	if err != nil {
		t.Fatal(err)
	}

	// httptest.Recorder записывает ответ
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(statusHandler)

	// Вызываем наш обработчик
	handler.ServeHTTP(rr, req)

	// Проверяем, что код ответа 200 OK
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler вернул неверный код: получили %v, хотели %v",
			status, http.StatusOK)
	}

	expected := `Go web App`
	if body := rr.Body.String(); strings.Contains(body, expected) {
		t.Errorf("handler вернул неожиданное тело: получили %v, не содержит %v",
			body, expected)
	}
}
