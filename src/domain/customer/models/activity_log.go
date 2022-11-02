package models

import (
	"time"
)

// Customer activity logs Model
type ActivityLog struct {
	ID          uint       `json:"id"`
	CustomerId  string     `json:"customer_id"`
	Description *string    `json:"description"`
	Properties  *string    `json:"properties"`
	Birthday    *time.Time `json:"time"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

// TableName gives table name of model
func (log ActivityLog) TableName() string {
	return "activity_logs"
}
