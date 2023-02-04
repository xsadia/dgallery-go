package main

import "github.com/xsadia/kgallery/pkg/server"

func main() {
	app := server.NewServer()

	app.ListenAndServe()
}
