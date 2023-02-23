package config

import (
	"os"
)

type Context struct {
	Env map[string]string
}

var (
	Ctx     *Context
	EnvKeys = []string{
		"POSTGRES_URL",
		"PGSQL_HOST",
		"PGSQL_USER",
		"PGSQL_PASSWORD",
		"PGSQL_DBNAME",
		"PGSQL_NAME",
		"DISCORD_CLIENT_ID",
		"DISCORD_SECRET",
	}
)

func Init() {
	envMap := make(map[string]string, len(EnvKeys))

	for _, v := range EnvKeys {
		envMap[v] = os.Getenv(v)
	}

	envMap["PORT"] = getPort()

	Ctx = &Context{
		Env: envMap,
	}
}

func getPort() string {
	port := os.Getenv("PORT")

	if []byte(port)[0] == ':' {
		return port
	}

	return string(append([]byte(":"), []byte(port)...))
}
