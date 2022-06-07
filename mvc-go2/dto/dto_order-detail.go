package dto

type OrderDetail struct {
	Amount  float32    `json:"total"`
	Price   float32    `json:"price"`
	Product ProductDto `json:"product"` //ES LO Q ESTÁ EN LA OTRA CLASE
}

type OrdersDetailsDto []OrdersDetailsDto //DEFINICION - DEFINO UN TIPO
