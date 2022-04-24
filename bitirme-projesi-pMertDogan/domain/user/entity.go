package user

import (
	"encoding/json"

	"gorm.io/gorm"
)

type Users []User

type User struct {
	gorm.Model
	UserName string 
	Email string 	`gorm:"unique"` //make sure email is unique 
	//binds are not needed cause they are not in frontline -> bind:"required,email". we use DTO for api requests
	Password string	
	IsAdmin  bool 	
	// FalseLoginCount int
	// ExpiresAt string
}

// fromJson users
func UnmarshalBooks(data []byte) (Users, error) {
	var r Users
	err := json.Unmarshal(data, &r)
	return r, err
}


//Users toJson
func (r *Users) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

//User toJson
func (r *User) Marshal() ([]byte, error) {
	return json.Marshal(r)
}


// func (r *User) FromRegisterRequestDTO(rUser RegisterRequestDTO) User{
// 	return User{UserName: r.UserName, Email: r.Email, Password: r.Password}
// }
