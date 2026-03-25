package actions

type UserSignUp struct {
	FullName        string `json:"full_name" validate:"required,min=2,max=50"`
	Email           string `json:"email" validate:"required,email"`
	Phone           string `json:"phone_number" validate:"kenyan_phone"`
	Password        string `json:"password" validate:"required,min=8"`
	ConfirmPassword string `json:"confirm_password" validate:"required,min=8"`
}

type UserSignIn struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}
