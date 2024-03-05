package model

type RolePermission struct {
	Id           int   `json:"id"`
	RoleId       int   `json:"role_id"`       //
	PermissionId int   `json:"permission_id"` //
	CreatedAt    int64 `json:"created_at"`    //
	UpdatedAt    int64 `json:"updated_at"`    //
}

func (RolePermission) TableName() string {
	return "role_permission"
}
