package models

type User struct {
	ID       uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name     string `json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Password []byte `json:"-"`
}
