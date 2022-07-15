package swagger

import "github.com/gofiber/swagger"

func InitSwagger() {
	swagger.New(swagger.Config{
		Title:        "Swagger API",
		DeepLinking:  false,
		DocExpansion: "none",
	})
}
