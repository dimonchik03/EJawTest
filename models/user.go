package models

type User struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"type:varchar(255);not null"`
	Role        string `gorm:"type:varchar(255);not null"`
	Login       string `gorm:"type:varchar(255);unique;not null"`
	Password    string `gorm:"type:text;not null"`
	PhoneNumber string `gorm:"type:varchar(12);not null"`
}
