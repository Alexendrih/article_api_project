package models

type User struct {
	ID         uint    `json:"user_id" gorm:"primaryKey"`
	UserNAme   string  `json:"login_username"`
	Password   string  `json:"login_password"`
	UserFnAme  string  `json:"user_fname"`
	UserLnAme  string  `json:"user_lname"`
	YearLevel  string  `json:"year_level"`
	Department string  `json:"department"`
	Email      string  `json:"email"`
	Phone      string  `json:"phone"`
	RoleId     int     `json:"user_role_id"`
	Userole    Userole `gorm:"foreignKey:RoleId"`
}
