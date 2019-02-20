package routes

import (
	"demo4Iris/controller"

	. "demo4Iris/common/config"
)

func InitRoute() {
	// user
	user := App.Party("user")
	user.Get("/", controller.GetUser)
}
