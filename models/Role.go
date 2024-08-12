package models

type Role struct {
	Id   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"size:100" json:"name"`
}
