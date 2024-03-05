package model

type Permission struct {
	Id        int    `json:"id"`         //
	Key       string `json:"key"`        //
	Type      string `json:"type"`       //
	Desc      string `json:"desc"`       //
	Attribute string `json:"attribute"`  //
	CreatedAt int64  `json:"created_at"` //
	UpdatedAt int64  `json:"updated_at"` //
}

func (Permission) TableName() string {
	return "permission"
}
