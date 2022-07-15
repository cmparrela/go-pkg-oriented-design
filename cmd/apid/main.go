package main

import (
	"github.com/gofiber/swagger"
	"go-pkg-oriented-design/internal/book"
	"go-pkg-oriented-design/internal/platform/config"
	"go-pkg-oriented-design/internal/platform/database"
	"go-pkg-oriented-design/internal/platform/http"
	"go-pkg-oriented-design/internal/platform/identifier"
	"go-pkg-oriented-design/internal/platform/validator"
	"go-pkg-oriented-design/internal/user"
	"log"

	validatorv10 "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func main() {
	swagger.New(swagger.Config{
		Title:        "Swagger API",
		DeepLinking:  false,
		DocExpansion: "none",
	})

	envConfig := config.SetupEnvFile()

	mongo := database.InitMongo(envConfig.MongoAddress, envConfig.DatabaseName)

	identifier := identifier.NewIdentifier()
	validator := validator.NewValidator(validatorv10.New())
	userRepository := user.NewRepository(mongo)
	userService := user.NewService(userRepository, validator, identifier)
	userHandler := user.NewHandler(userService)

	bookRepository := book.NewRepository(mongo)
	bookService := book.NewService(bookRepository, validator, identifier)
	bookHandler := book.NewHandler(bookService)

	app := fiber.New(fiber.Config{
		ErrorHandler: http.ErrorHandler,
	})

	http.SetupRoutes(
		app,
		userHandler,
		bookHandler,
	)

	if err := app.Listen(":5001"); err != nil {
		log.Fatalf("listen: %s", err)
	}
}
