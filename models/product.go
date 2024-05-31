package models

type Product struct {
	ID           uint    `gorm:"primaryKey"`
	Name         string  `gorm:"type:varchar(255);not null"`
	Description  string  `gorm:"type:text;not null"`
	SerialNumber string  `gorm:"type:varchar(12);not null"`
	SellerID     uint    `gorm:"not null"`
	Seller       Contact `gorm:"foreignKey:SellerID;references:ID"`
}
