package models

import "time"

type Article struct {
	ID         uint   `json:"article_id" gorm:"primaryKey"`
	Title      string `json:"title"`
	Body       string `json:"body"`
	Footer     string `json:"footer"`
	CreatedAt  time.Time
	Remarks    string `json:"article_remarks"`
	Status     string `json:"date_approved"`
	ApprovedBy string `json:"approved_by"`
	UserId     int    `json:"user_id"`
	User       User   `gorm:"foreignKey:UserId"`
}
