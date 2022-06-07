package productController

import (
	"fmt"
	"mvc-go/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProductByEan(c *gin.Context) {
	c.JSON(http.StatusOK, "Buscando: "+c.Param("product_id"))
}

func GetProducts(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}

//--
func GetProductByName(c *gin.Context) {

	name := c.Param("KeyPro")

	var productDto dto.ProductDto

	productDto, err := service.product_service.GetProductByName(name)

	if err != nil {

		c.JSON(err.Status(), err)
		return

	}

	c.JSON(http.StatusOK, productDto)

}

func GetProductByCat(c *gin.Context) {

	cat := c.Param("KeyCat")

	productsDto, err := service.ProductService.GetProductByCat(cat)
	fmt.Println("todas las consultas ok")
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, productsDto)
}

func GetProductAll(c *gin.Context) {
	productsDto, err := service.ProductService.GetProductAll()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, productsDto)
}
