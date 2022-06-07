package dto

type OrderDto struct {
	OrderDate      string           `json:"order_date"`
	Total          float32          `json:"total"`
	Status         string           `json:"status"`
	User           UserDto          `json:"user"`
	OrderDetailDto OrdersDetailsDto `json:"detail"`
}

type OrdersDto []OrderDto
