package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mhd7966/Vafa/controllers"
	"github.com/mhd7966/Vafa/log"
)

func UserRouter(app fiber.Router) {

	api := app.Group("/user")

	api.Get("/", controllers.GetUsers)
	api.Get("/:id", controllers.GetUser)
	api.Post("/", controllers.NewUser)
	api.Post("/:id", controllers.CancelUser)
	api.Put("/:id", controllers.UpdateUser)
	api.Delete("/:id", controllers.DeleteUser)

	log.Log.Info("User routes created :)")

}
