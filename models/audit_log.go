package models

type Audit struct {
	ID       uint   `json:"audit_log_id" gorm:"primaryKey"`
	Register string `json:"date_register"`
	Indate   string `json:"login_date"`
	Outdate  string `json:"logut_date"`
	Userstat string `json:"user_status"`
	UserId   int    `json:"user_id"`
	User     User   `gorm:"foreignKey:UserId"`
}
