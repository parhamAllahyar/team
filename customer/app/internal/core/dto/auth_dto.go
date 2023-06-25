package dto

type SignInRequest struct {
	Password string `validate:"required"`
	Email    string `validate:"required,email"`
}

type SignUpRequest struct {
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
	Password  string `validate:"required"`
	Email     string `validate:"required,email"`
	Phone     string `validate:"required"`
}

type AuthResponse struct {
	AccessToken string
}

type ForgetPasswordRequest struct{}

type UpdatePasswordRequest struct {
	AccessToken string
	OldPassword string
	NewPassword string
}
