package dto

type Video struct {
	Title       string `json:"title" binding:"required,min=2,max=50"`
	Description string `json:"description"`
}
