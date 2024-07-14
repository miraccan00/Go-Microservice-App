package models

import "time"

type Step struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id"`
	Count     int       `json:"count"`
	Timestamp time.Time `json:"timestamp"`
}
