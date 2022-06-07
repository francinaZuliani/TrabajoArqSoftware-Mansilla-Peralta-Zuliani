package model

type Category struct {
	NameCategory        string `gorm:"type:varchar(350);not null"`
	DescriptionCategory string `gorm:"type:varchar(850);not null"`
	//Id                  int    `gorm:"primaryKey"`
}
