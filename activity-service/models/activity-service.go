package models

import "time"

type Activity struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id"`
	Type      string    `json:"type"`
	Duration  float64   `json:"duration"` // Duration in hours
	Timestamp time.Time `json:"timestamp"`
}
