package dto

type UserDto struct {
	Name        string `json:"name"`
	LastName    string `json:"last_name"`
	UserName    string `json:"user_name"`
	Address     string `json:"address"`
	PhoneNumber int    `json:"phone_number"`
	Password    string `json:"password"`
	Id          int    `json:"id"`
}

type UsersDto []UserDto
