package models

type OrderProduct struct {
	ID        uint    `gorm:"primaryKey"`
	OrderID   uint    `gorm:"not null"`
	ProductID uint    `gorm:"not null"`
	Amount    uint    `gorm:"not null"`
	Order     Order   `gorm:"foreignKey:OrderID;references:ID"`
	Product   Product `gorm:"foreignKey:ProductID;references:ID"`
}
