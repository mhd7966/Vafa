package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mhd7966/Vafa/controllers"
	"github.com/mhd7966/Vafa/log"
)

func UserRouter(app fiber.Router) {

	api := app.Group("/users")

	api.Get("/", controllers.GetUsers)
	api.Get("/:id", controllers.GetUser)
	api.Post("", controllers.NewUser)
	api.Put("/:id/cancel", controllers.CancelUser)
	api.Delete("/:id", controllers.DeleteUser)

	log.Log.Info("User routes created :)")

}
