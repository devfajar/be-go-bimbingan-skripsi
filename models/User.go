package models

import "time"

type User struct {
	ID        uint      `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `json:"name"`
	Email     string    `gorm:"unique" json:"email"`
	Password  []byte    `json:"-"`
	RoleID    uint      `json:"role_id"`
	Role      Role      `gorm:"foreignKey:RoleID" json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
