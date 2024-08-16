package models

import "time"

type Role struct {
	Id        uint      `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"size:100" json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
