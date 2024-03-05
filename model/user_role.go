package model

type UserRole struct {
	Id        int    `json:"id"`
	UserId    int    `json:"user_id"`    //
	RoleId    int    `json:"role_id"`    //
	RoleName  string `json:"role_name"`  //记录角色名称，检查查表次数
	CreatedAt int64  `json:"created_at"` //
	UpdatedAt int64  `json:"updated_at"` //
}

func (UserRole) TableName() string {
	return "user_role"
}
