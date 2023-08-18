package contracts

import (
	"testproject/internal/services/videoService/dto"
)

type IVideoService interface {
	All() ([]dto.Video, error)
	Creat(video dto.Video) (*dto.Video, error)
	GetUserVideos(userId uint) ([]dto.Video, error)
	GetFavoriteVideos(userId uint) ([]dto.Video, error)
}
