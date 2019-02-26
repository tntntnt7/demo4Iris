package routes

import (
	. "github.com/tntntnt7/demo4Iris/common/config"
	. "github.com/tntntnt7/demo4Iris/controller"
)

func InitRoute() {
	App.Post("/login", Login)
	// user
	user := App.Party("user")
	user.Get("/{id}", GetUser)
	user.Get("/all", GetAllUsers)
	user.Post("/", Register)
	user.Put("/", UpdateUser)
	user.Delete("/{id}", DeleteUser)
}
