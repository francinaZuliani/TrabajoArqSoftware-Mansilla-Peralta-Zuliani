package model

type Product struct {
	Id            int     `gorm:"primaryKey"`
	Name          string  `gorm:"type:varchar(350);not null"`
	UniversalCode string  `gorm:"type:varchar(12);not null"`
	Picture       string  `gorm:"type:varchar(250);not null"`
	Price         float32 `gorm:"type:float;not null"`
	Description   string  `gorm:"type:varchar(280);not null;unique"`
	Stock         int     `gorm:"type:int;not null"`
	Reviews       string  `gorm:"type:varchar(280);not null"`
}

type ProductsDto []ProductsDto
