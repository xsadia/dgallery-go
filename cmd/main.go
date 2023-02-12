package main

import (
	"github.com/xsadia/kgallery/config"
	"github.com/xsadia/kgallery/pkg/server"
)

func main() {
	app := server.NewServer()
	config.Init("../.env")

	app.ListenAndServe()
}
