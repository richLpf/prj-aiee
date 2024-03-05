package model

type User struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Name      string `json:"name"`
	Status    int64  `json:"status"`
	Age       int64  `json:"age"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

func (User) TableName() string {
	return "user"
}
