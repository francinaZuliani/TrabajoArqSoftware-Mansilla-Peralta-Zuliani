package product

import (
	"mvc-go/model"

	//"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

//var Db *gorm.DB

func GetProductByName(name string) []model.Product {
	var products []model.Product
	Db.Where("name LIKE ?", name+"%").Find(&products)
	log.Debug("products: ", products)

	return products
}

func GetProductByUniversalCode(UniversalCode string) model.Product { //...
	var product model.Product

	Db.Where("universal_code = ?", UniversalCode).First(&product)
	log.Debug("Product: ", product)

	return product
}

func GetProductByCategory(namecategory string) []model.Product { //

	var products []model.Product

	Db.Where("name LIKE ?", namecategory+"%").Find(&products)

	log.Debug("products: ", products)

	return products

}

func GetProductAll() []model.Product {
	var products []model.Product

	Db.Find(&products)

	log.Debug("products: ", products)

	return products
}
