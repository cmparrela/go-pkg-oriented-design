package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/swagger"
	_ "go-pkg-oriented-design/docs"
	"go-pkg-oriented-design/internal/book"
	"go-pkg-oriented-design/internal/user"
)

func SetupRoutes(
	app fiber.Router, userHandler user.Handler, bookHandler book.Handler) {
	app.Get("/", monitor.New())

	app.Get("/docs/*", swagger.HandlerDefault)

	app.Get("/users", userHandler.List)
	app.Get("/users/:id", userHandler.Find)
	app.Put("/users/:id", userHandler.Update)
	app.Post("/users", userHandler.Create)
	app.Delete("/users/:id", userHandler.Delete)

	app.Get("/books", bookHandler.List)
	app.Get("/books/:id", bookHandler.Find)
	app.Put("/books/:id", bookHandler.Update)
	app.Post("/books", bookHandler.Create)
	app.Delete("/books/:id", bookHandler.Delete)
}
