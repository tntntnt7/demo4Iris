package routes

import (
	"github.com/tntntnt7/demo4Iris/controller"

	. "github.com/tntntnt7/demo4Iris/common/config"
)

func InitRoute() {
	App.Post("/login", controller.Login)
	// user
	user := App.Party("user")
	user.Get("/{id}", controller.GetUser)
	user.Get("/all", controller.GetAllUsers)
	user.Post("/", controller.Register)
	user.Put("/", controller.UpdateUser)
	user.Delete("/{id}", controller.DeleteUser)
}
