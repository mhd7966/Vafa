package routes

import (

	"github.com/gofiber/fiber/v2"
	"github.com/mhd7966/Vafa/log"
)

func MainRouter(app fiber.Router) {
	api := app.Group("/v0", gopkgs.Auth)

	LogRouter(api)
	HooRouter(api)
	AlertRouter(api)
	UserRouter(api)
	AlertGroupRouter(api)

	log.Log.Info("All routes created successfully :)")

}
