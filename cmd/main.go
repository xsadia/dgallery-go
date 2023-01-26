package main

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	HTTP *fiber.App
}

func NewServer() *Server {
	srv := &Server{
		HTTP: fiber.New(fiber.Config{
			CaseSensitive: true,
			StrictRouting: true,
		}),
	}

	srv.routes()
	return srv
}

func (s *Server) routes() {
	s.HTTP.Get("/ping", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(&fiber.Map{
			"data": "pong",
		})
	})
}

func (s *Server) listenAndServe() {
	if err := s.HTTP.Listen("8080"); err != nil {
		log.Fatalf("[Error]: failed to serve on port 8080, %s\n", err.Error())
	}
}

func main() {
	app := NewServer()

	app.listenAndServe()
}
