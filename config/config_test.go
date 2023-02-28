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

		expected := make(map[string]string, len(EnvKeys)+1)

		expected["PORT"] = ":8081"

		for _, v := range EnvKeys {
			expected[v] = v
		}

		assert.Equal(expected, Ctx.Env)
	})
}
