package permission

import "prj-aiee/router"

func init() {
	group := router.NewGroup("permission")
	group.NewRouter("/create", CreatePermission)
	group.NewRouter("/getList", GetPermissionList)
	group.NewRouter("/get", GetPermission)
	group.NewRouter("/update", UpdatePermission)
	group.NewRouter("/delete", DeletePermission)
	group.Register()
}
