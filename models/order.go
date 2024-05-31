package models

type Order struct {
	ID   uint   `gorm:"primaryKey"`
	Date string `gorm:"type:date;not null"`
}

type OrderStruct struct {
	Order Order
	Items []OrderProduct
}
