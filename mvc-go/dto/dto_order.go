package dto

type OrderDto struct {
	Id               int              `json:"id"`
	OrderDate        string           `json:"order_date"`
	Total            float32          `json:"total"`
	Status           string           `json:"status"`
	User             UserDto          `json:"user"`
	OrdersDetailsDto OrdersDetailsDto `json:"detail"`
}

type OrdersDto []OrderDto
