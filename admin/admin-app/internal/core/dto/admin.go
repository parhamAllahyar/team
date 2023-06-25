package dto

// import "github.com/google/uuid"

type UpdatePasswordRequest struct {
	AccessToken string
	NewPassword string `json:"new_password"`
	OldPassword string `json:"old_password"`
}

type UpdateRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type UpdateResponse struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type UpdatePoneRequest struct {
	AccessToken string
	Phone       string `json:"phone"`
}

type InviteLinkRequest struct {
	AccessToken string
	Email       string `json:"email"`
}
