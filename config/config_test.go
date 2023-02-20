package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	t.Run("Loads env variables", func(t *testing.T) {

		assert := assert.New(t)

		assert.Nil(Ctx, "expected Ctx to not be instantiated at the beginning")

		os.Setenv("PORT", "8081")

		for _, v := range EnvKeys {
			os.Setenv(v, v)
		}

		Init()

		assert.NotNil(Ctx.Env, "expected Ctx to be instantiated after Init is called, got nil")

		expected := map[string]string{
			"PGSQL_DBNAME":   "PGSQL_DBNAME",
			"PGSQL_HOST":     "PGSQL_HOST",
			"PGSQL_NAME":     "PGSQL_NAME",
			"PGSQL_PASSWORD": "PGSQL_PASSWORD",
			"PGSQL_USER":     "PGSQL_USER",
			"PORT":           ":8081",
			"POSTGRES_URL":   "POSTGRES_URL"}

		assert.Equal(expected, Ctx.Env)
	})
}
