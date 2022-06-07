package model

type OrderDto struct {
	Id        int     `gorm:"primaryKey"`
	OrderDate string  `gorm:"type:varchar(350);not null"`
	Total     float32 `gorm:"type:float;not null"`
	Status    string  `gorm:"type:varchar(350);not null"`
}

type OrdersDto []OrderDto
