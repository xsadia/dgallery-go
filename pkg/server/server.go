package server

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/xsadia/kgallery/config"
	"github.com/xsadia/kgallery/pkg/storage"
)

type Server struct {
	HTTP    *fiber.App
	Storage *storage.Storage
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
		return c.Status(http.StatusOK).JSON(map[string]string{
			"data": "pong",
		})
	})
}

func (s *Server) ListenAndServe() {
	if err := s.HTTP.Listen(config.Ctx.Env["PORT"]); err != nil {
		log.Fatalf("[Error]: failed to serve on port %s, %s\n", config.Ctx.Env["PORT"], err.Error())
	}
}

func (s *Server) Close() {
	if err := s.HTTP.Shutdown(); err != nil {
		log.Fatalf("[Error]: %s", err.Error())
	}
}
