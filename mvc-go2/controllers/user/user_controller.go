package userController

import (
	"mvc-go/dto"
	"mvc-go/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	log "github.com/sirupsen/logrus"
)

var jwtKey = []byte("secret_key")

func GetUserById(c *gin.Context) { //recibe con id
	log.Debug("User id: " + c.Param("id"))

	var userDto dto.UserDto
	userDto.Name = "Fran"
	userDto.LastName = "Zuliani"

	c.JSON(http.StatusOK, userDto)
}

func GetUsers(c *gin.Context) {

	c.JSON(http.StatusOK, "Users:")
}

func UserInsert(c *gin.Context) { //recibe un parámetro json
	var userDto dto.UserDto
	err := c.BindJSON(&userDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusCreated, userDto)
}

func Login(c *gin.Context) {
	var loginDto dto.LoginDto
	err := c.BindJSON(&loginDto)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	tokenDto, er := services.UserService.Login(loginDto)

	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	tkn, err := jwt.Parse(tokenDto.Token, func(t *jwt.Token) (interface{}, error) { return jwtKey, nil })

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.JSON(http.StatusUnauthorized, "Invalid Signature")
			return
		}
		c.JSON(http.StatusBadRequest, "Bad Request")
		return
	}

	if !tkn.Valid {
		c.JSON(http.StatusUnauthorized, "Invalid Token")
		return
	}
	c.JSON(http.StatusCreated, tokenDto)

}
