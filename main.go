package main

import (
	"github.com/calvinchengx/fiber-go-pg/config"
	"github.com/go-pg/pg/v9"
	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()
	db := config.GetConnection()
	rs := &RouteServices{db, app}
	rs.setupV1Routes()
	app.Listen(3000)
}

// RouteServices lets us bind specific services when setting up routes
type RouteServices struct {
	db  *pg.DB
	app *fiber.App
}

func (s *RouteServices) setupV1Routes() {
	s.app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello World!")
	})
	v1Router := s.app.Group("/v1")
	v1Router.Use()
}
