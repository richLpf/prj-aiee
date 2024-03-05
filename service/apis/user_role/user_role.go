package user_role

import "prj-aiee/router"

func init() {
	group := router.NewGroup("user_role")
	group.NewRouter("/create", CreateUserRole)
	group.NewRouter("/getList", GetUserRoleList)
	group.NewRouter("/get", GetUserRole)
	group.NewRouter("/update", UpdateUserRole)
	group.NewRouter("/delete", DeleteUserRole)
	group.Register()
}
