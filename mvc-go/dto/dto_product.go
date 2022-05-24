package dto

type ProductDto struct {
	Name          string      `json:"name"`
	UniversalCode string      `json:"universal_code"`
	Picture       string      `json:"picture_url"`
	Price         float32     `json:"base_price"`
	Description   string      `json:"description"`
	Stock         int         `json:"stock"`
	Reviews       string      `json:"reviews"`
	Id            int         `json:"id"`
	Category      CategoryDto `json:"category"`
}

type ProductsDto []ProductDto
