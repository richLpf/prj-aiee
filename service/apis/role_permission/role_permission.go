package role_permission

import "prj-aiee/router"

func init() {
	group := router.NewGroup("role_permission")
	group.NewRouter("/create", CreateRolePermission)
	group.NewRouter("/getList", GetRolePermissionList)
	group.NewRouter("/get", GetRolePermission)
	group.NewRouter("/update", UpdateRolePermission)
	group.NewRouter("/delete", DeleteRolePermission)
	group.Register()
}
