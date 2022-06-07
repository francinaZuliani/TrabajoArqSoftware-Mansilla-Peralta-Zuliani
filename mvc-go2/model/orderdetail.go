package model

type OrderDetail struct {
	Id     int     `gorm:"primaryKey"`
	Amount float32 `gorm:"type:float;not null"`
	Price  float32 `gorm:"type:float;not null"`
}

type OrdersDetailsDto []OrdersDetailsDto //DEFINICION - DEFINO UN TIPO
