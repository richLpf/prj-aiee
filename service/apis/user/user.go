package user

import "prj-aiee/router"

func init() {
	group := router.NewGroup("user")
	group.NewRouter("/create", CreateUser)
	group.NewRouter("/login", Login)
	group.NewRouter("/getList", GetUserList)
	group.NewRouter("/get", GetUser)
	group.NewRouter("/update", UpdateUser)
	group.NewRouter("/delete", DeleteUser)
	group.Register()
}
