package dto

type RegisterRequest struct {
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
	Email    string `json:"email"`
}
