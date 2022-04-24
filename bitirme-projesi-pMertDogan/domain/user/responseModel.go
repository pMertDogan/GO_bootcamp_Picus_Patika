package user

type ResponseModel struct{
	ErrMsg string `json:"errMsg,omitempty"` //if errMsg is empty, it means there is no error. With the help of the omitempty, we can omit the field from the response
	Err string `json:"errBody,omitempty"` 
	ResponseCode int `json:"responseCode"`
	AccesToken string `json:"accessToken,omitempty"`
	RefreshToken string	`json:"refreshToken,omitempty"`
	UserData UserToReponseDTO `json:"userData,omitempty"`
}

//This one is used to return the response to the client
type UserToReponseDTO struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

