package main

import (
	"github.com/xsadia/kgallery/config"
	"github.com/xsadia/kgallery/pkg/server"
)

func main() {
	config.Init()
	app := server.NewServer()

	app.ListenAndServe()
}
