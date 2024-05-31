package models

type Contact struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"type:varchar(255);not null"`
	Role        string `gorm:"type:varchar(255);not null"`
	PhoneNumber string `gorm:"type:varchar(12);not null"`
}
