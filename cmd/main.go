package main

import (
	"github.com/xsadia/kgallery/config"
	"github.com/xsadia/kgallery/pkg/server"
	"github.com/xsadia/kgallery/pkg/storage"
)

func main() {
	config.Init("../.env")
	app := server.NewServer()
	app.Storage = storage.NewStorage(config.Ctx)

	app.ListenAndServe()
}
