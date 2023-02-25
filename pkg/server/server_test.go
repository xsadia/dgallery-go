package server

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xsadia/kgallery/config"
)

func TestMain(m *testing.M) {
	exitVal := m.Run()

	os.Exit(exitVal)
}

func TestRouter(t *testing.T) {
	t.Run("returns pong", func(t *testing.T) {

		os.Setenv("PORT", "8081")

		for _, v := range config.EnvKeys {
			os.Setenv(v, v)
		}

		config.Init()

		assert := assert.New(t)

		app := NewServer()

		req := httptest.NewRequest(http.MethodGet, "/ping", nil)

		resp, err := app.HTTP.Test(req, 1)

		assert.Nil(err)

		var m map[string]string
		body, err := io.ReadAll(resp.Body)

		assert.Nil(err)

		json.Unmarshal(body, &m)

		assert.Equal(resp.StatusCode, http.StatusOK)

		assert.Equal(m["data"], "pong")
	})
}
