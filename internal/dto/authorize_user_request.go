package dto

type AuthorizeUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
