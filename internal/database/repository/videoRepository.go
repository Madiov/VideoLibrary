package repository

import (
	"gorm.io/gorm"
	"testproject/internal/services/videoService/dto"
)

type Video struct {
	gorm.Model
	Title       string
	Description string
	User        User
	USerID      string
}
type FavoriteVideo struct {
	UserID  uint
	VideoID uint
}
type VideoRepository struct {
	db *gorm.DB
}

func NewVideoRepository(db *gorm.DB) *VideoRepository {
	return &VideoRepository{db: db}
}

func (v VideoRepository) GetFavoriteVideos(userId uint) ([]dto.Video, error) {
	//TODO implement me
	panic("implement me")
}
