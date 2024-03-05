package model

type Role struct {
	Id        int    `json:"id"`
	Key       string `json:"key"`        //
	Name      string `json:"name"`       //
	Desc      string `json:"desc"`       //
	Status    int    `json:"status"`     //
	CreatedAt int64  `json:"created_at"` //
	UpdatedAt int64  `json:"updated_at"` //
}

func (Role) TableName() string {
	return "role"
}
