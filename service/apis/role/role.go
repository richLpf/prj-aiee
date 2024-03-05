package role

import "prj-aiee/router"

func init() {
	group := router.NewGroup("role")
	group.NewRouter("/create", CreateRole)
	group.NewRouter("/getList", GetRoleList)
	group.NewRouter("/get", GetRole)
	group.NewRouter("/update", UpdateRole)
	group.NewRouter("/delete", DeleteRole)
	group.Register()
}
