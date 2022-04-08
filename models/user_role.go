package models

type Userole struct {
	ID      uint   `json:"user_ role_id" gorm:"primaryKey"`
	Role    string `json:"role"`
	Remarks string `json:"remarks"`
}
