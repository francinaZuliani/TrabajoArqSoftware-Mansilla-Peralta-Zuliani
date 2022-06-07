package services

import (
	userCliente "mvc-go/clients/user"
	"mvc-go/dto"
	"mvc-go/model"
	e "mvc-go/utils/errors"

	"github.com/golang-jwt/jwt"
)

type userService struct{}

var jwtKey = []byte("secret_key")

type userServiceInterface interface {
	GetUserById(id int) (dto.UserDto, e.ApiError)
	GetUsers() (dto.UsersDto, e.ApiError)
	InsertUser(userDto dto.UserDto) (dto.UserDto, e.ApiError)
	//Login(loginDto dto.LoginDto) (dto.LoginDto, e.ApiError)
	Login(loginDto dto.LoginDto) (dto.Token, e.ApiError)
}

var (
	UserService userServiceInterface
)

func init() {
	UserService = &userService{}
}

func (s *userService) GetUserById(id int) (dto.UserDto, e.ApiError) {

	var user model.User = userCliente.GetUserById(id)
	var userDto dto.UserDto

	if user.Id == 0 {
		return userDto, e.NewBadRequestApiError("user not found")
	}
	userDto.Name = user.Name
	userDto.LastName = user.LastName
	//userDto.UserName = user.UserName
	//userDto.Id = user.Id
	return userDto, nil
}

func (s *userService) GetUsers() (dto.UsersDto, e.ApiError) {

	var users model.Users = userCliente.GetUsers()
	var usersDto dto.UsersDto

	for _, user := range users {
		var userDto dto.UserDto
		userDto.Name = user.Name
		userDto.LastName = user.LastName
		//userDto.UserName = user.Name
		//userDto.Id = user.Id

		usersDto = append(usersDto, userDto)
	}

	return usersDto, nil
}

func (s *userService) InsertUser(userDto dto.UserDto) (dto.UserDto, e.ApiError) {

	var user model.User

	user.Name = userDto.Name
	user.LastName = userDto.LastName
	//user.UserName = userDto.UserName
	//user.Password = userDto.Password

	user = userCliente.InsertUser(user)

	userDto.Id_user = user.Id

	return userDto, nil
}

func (s *userService) Login(loginDto dto.LoginDto) (dto.Token, e.ApiError) {
	var user model.User = userCliente.GetUserByName(loginDto.Usuario)

	var tokenDto dto.Token

	if user.Id == 0 {
		return tokenDto, e.NewBadRequestApiError("User not found")
	}

	if user.Password == loginDto.Password {
		token := jwt.New(jwt.SigningMethodHS256)
		tokenString, _ := token.SignedString(jwtKey)
		tokenDto.Token = tokenString
		tokenDto.Id_user = user.Id
	}
	return tokenDto, nil
}
