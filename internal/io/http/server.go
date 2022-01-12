package http

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/khorevaa/r2gitsync/internal/bl"
	"github.com/khorevaa/r2gitsync/internal/io/http/routes"
)

type Server struct {
	bl *bl.BL
}

func New(bl *bl.BL) *Server {
	return &Server{
		bl: bl,
	}
}

func (s *Server) Run(addr string) {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(requestid.New())
	routes.SetUpRoutes(app, s.bl)
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	log.Fatal(app.Listen(addr))
}
