package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserCanCheckIfPlayer1HasWon(t *testing.T) {

	jsonData := []byte(`
	{
		"matrix": [
			[1, 1, 1],
			[2, 1, 0],
			[2, 1, 0]
		],
		"player": 1
	}`)

	req, err := http.NewRequest("POST", "/games/is-won", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(IsWonHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"isWon": true}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestUserCanCheckIfPlayer2HasWon(t *testing.T) {

	jsonData := []byte(`
	{
		"matrix": [
			[1, 1, 0],
			[2, 0, 2],
			[2, 2, 2]
		],
		"player": 2
	}`)

	req, err := http.NewRequest("POST", "/games/is-won", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(IsWonHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"isWon": true}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
