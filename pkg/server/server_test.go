package server

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	exitVal := m.Run()

	os.Exit(exitVal)
}

func TestRouter(t *testing.T) {
	t.Run("returns pong", func(t *testing.T) {

		expected := "pong"

		app := NewServer()

		req := httptest.NewRequest("GET", "/ping", nil)

		resp, err := app.HTTP.Test(req, 1)

		if err != nil {
			t.Errorf("Expected error to be nil, got %s", err.Error())
		}

		var m map[string]string
		body, err := io.ReadAll(resp.Body)

		if err != nil {
			t.Errorf("Expected error to be nil, got %s", err.Error())
		}

		json.Unmarshal(body, &m)

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected %d, got %d", http.StatusOK, resp.StatusCode)
		}

		if m["data"] != expected {
			t.Errorf("Expected %s, got %s", expected, m["message"])
		}
	})
}
