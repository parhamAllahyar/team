package dto

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	AccessToken string
}

type ForgetPasswordRequest struct {
	Email string `json:"email"`
}

type SignUpRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Link      string `uri:"link"`
	Password  string `json:"password"`
}

type GenerateLinkRequest struct {
	AccessToken string
	Email       string `json:"email"`
}
