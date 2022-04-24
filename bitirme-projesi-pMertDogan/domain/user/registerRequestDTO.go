package user

//this is our model for login request
type RegisterRequestDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	Name string `json:"name" binding:"required,min=3,alphanum"`
}

//gin is using validator paca to validate our request and we are adding validation check with binding tag
// https://blog.logrocket.com/gin-binding-in-go-a-tutorial-with-examples/
// https://github.com/go-playground/validator

/*


Tag	Description	Usage Example
uppercase	Accepts only uppercase letters	binding:"uppercase"
lowercase	Accepts only lowercase letters	binding:"lowercase"
contains	Accepts only strings that contain a specific string segment.	binding:"contains=key"
alphanum	Accepts only alphanumeric characters (English letters and numerical digits). Rejects strings that contain special characters.	binding:"alphanum"
alpha	Accepts only English letters	binding:"alpha"
endswith	Accepts only strings that end with a specific sequence of characters	binding:"endswith=."

*/
