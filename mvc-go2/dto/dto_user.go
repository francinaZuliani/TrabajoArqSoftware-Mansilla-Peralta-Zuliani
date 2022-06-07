package dto

type UserDto struct {
	Id_user     int    `json:"Id_user"`
	Name        string `json:"name"`
	LastName    string `json:"last_name"`
	Address     string `json:"address"`
	PhoneNumber int    `json:"phone_number"`
}

type UsersDto []UserDto

//NO SE RELACIONA NINUGNA CLASE
