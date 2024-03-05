package model

// CommonFields represents common fields for models.
type CommonFields struct {
	CreatedAt int64 `json:"created_at"` //
	UpdatedAt int64 `json:"updated_at"` //
}

type CreatedUser struct {
	CreatedUser string `json:"CreatedUser"`
}
