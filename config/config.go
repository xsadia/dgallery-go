package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Context struct {
	Env map[string]string
}

var (
	Ctx     *Context
	EnvKeys = []string{"POSTGRES_URL"}
)

func Init(envPath string) {
	if err := godotenv.Load(envPath); err != nil {
		log.Fatalf("[Error]: Error loading .env file, %v", err)
	}

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
