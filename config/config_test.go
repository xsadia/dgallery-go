package config

import (
	"testing"
)

func TestConfig(t *testing.T) {
	t.Run("Loads env variables", func(t *testing.T) {
		if Ctx != nil {
			t.Errorf("expected Ctx to not be instantiated at the beginning, got %v", Ctx)
		}

		Init("../.env.ci")

		if Ctx.Env == nil {
			t.Errorf("expected Ctx to be instantiated after Init is called, got nil")
		}

		for _, v := range EnvKeys {
			_, ok := Ctx.Env[v]

			if !ok {
				t.Errorf("Expected key: %v to be set", v)
			}
		}
	})
}
