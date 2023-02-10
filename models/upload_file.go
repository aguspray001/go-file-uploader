package models

import (
	"time"
)

type UploadFile struct {
	ID        int64     `gorm:"primaryKey" json:"id"`
	FileName  string    `json:"file_name"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"updated_at"`
}
